package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func GetPostgreSqlDB(dbHost string, dbPort int, dbUser, dbPassword, dbName string) (*sqlx.DB, error) {
	var sqlDB *sql.DB

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open PostgreSQL DB")
	}

	return sqlx.NewDb(sqlDB, "postgres"), nil
}
