package database

import (
	"database/sql"
	"fmt"
	"taal-client/settings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func GetPostgreSqlDB() (*sqlx.DB, error) {
	var sqlDB *sql.DB

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		settings.Get("dbHost"),
		settings.GetInt("dbPort"),
		settings.Get("dbUsername"),
		settings.Get("dbPassword"),
		settings.Get("dbName"),
	)

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open PostgreSQL DB")
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping PostgreSQL DB")
	}

	return sqlx.NewDb(sqlDB, "postgres"), nil
}
