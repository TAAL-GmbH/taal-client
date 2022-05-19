package repository

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrations embed.FS

const (
	dbFolder         = "localdata"
	sqLiteDBFilename = "db"
)

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
	if err := os.Mkdir(dbFolder, 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	sqliteDb, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s", dbFolder, sqLiteDBFilename))
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
