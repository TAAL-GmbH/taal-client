package database

import (
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	//go:embed migrations
	migrations embed.FS
)

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
