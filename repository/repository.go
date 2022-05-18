package repository

import (
	"context"
	"database/sql"
	"embed"
	"net/http"
	"taal-client/service"

	"github.com/golang-migrate/migrate/v4"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrations embed.FS

func NewDB(dbPath string) (*sqlx.DB, error) {
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	db := sqlx.NewDb(sqliteDb, "sqlite3")

	return db, nil
}

func RunMigrations(db *sqlx.DB) error {
	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return errors.Wrap(err, "invalid source instance")
	}
	targetInstance, err := sqlite.WithInstance(db.DB, new(sqlite.Config))
	if err != nil {
		return errors.Wrap(err, "invalid target sqlite instance")
	}
	m, err := migrate.NewWithInstance(
		"httpfs", sourceInstance, "sqlite", targetInstance)
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
	query := `INSERT INTO keys (api_key, private_key, public_key, address) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, key.ApiKey, key.PrivateKey, key.PublicKey, key.Adress)

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
