package repository

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"taal-client/service"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrations embed.FS

const sqLiteDBPath = "./localdata/db"

func GetPostgreSqlDB(host string, port int, username string, password string, dbName string) (*sqlx.DB, error) {
	var sqlDB *sql.DB

	connectionString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", port, host, username, password, dbName)

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open PostgreSQL DB")
	}

	return sqlx.NewDb(sqlDB, "postgres"), nil
}

func GetSQLiteDB() (*sqlx.DB, error) {
	sqliteDb, err := sql.Open("sqlite3", sqLiteDBPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	return sqlx.NewDb(sqliteDb, "sqlite3"), nil
}

func RunMigrationsSQLite(db *sqlx.DB) error {
	targetInstance, err := sqlite.WithInstance(db.DB, new(sqlite.Config))
	if err != nil {
		return errors.Wrap(err, "invalid target sqlite instance")
	}
	return RunMigrations(targetInstance, "sqlite")
}

func RunMigrationsPostgreSQL(db *sqlx.DB) error {
	// config := new(postgres.Config)
	targetInstance, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "invalid target postgres instance")
	}
	return RunMigrations(targetInstance, "postgres")
}

func RunMigrations(driver database.Driver, databaseName string) error {
	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return errors.Wrap(err, "invalid source instance")
	}

	m, err := migrate.NewWithInstance("httpfs", sourceInstance, databaseName, driver)
	if err != nil {
		return errors.Wrap(err, "failed to initialize migrate instance")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return sourceInstance.Close()
}

type Repository struct {
	db sqlx.DB
}

func NewRepository(db sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) InsertKey(ctx context.Context, key service.Key) error {
	query := `INSERT INTO keys (api_key, private_key, public_key, address) VALUES ($1, $2, $3, $4);`
	_, err := r.db.ExecContext(ctx, query, key.ApiKey, key.PrivateKey, key.PublicKey, key.Address)

	return err
}

func (r Repository) GetKey(ctx context.Context, apiKey string) (service.Key, error) {
	query := `SELECT * FROM keys WHERE api_key = ? LIMIT 1`

	key := service.Key{}

	err := r.db.Get(&key, query, apiKey)
	if err != nil {
		return service.Key{}, err
	}

	return key, nil
}
