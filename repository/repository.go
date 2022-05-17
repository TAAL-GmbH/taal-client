package repository

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

//go:embed migrations
var migrations embed.FS

func NewDB(dbPath string) (*sql.DB, error) {
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	return sqliteDb, nil
}

func RunMigrations(db *sql.DB) error {
	sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	if err != nil {
		return errors.Wrap(err, "invalid source instance")
	}
	targetInstance, err := sqlite.WithInstance(db, new(sqlite.Config))
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
