package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
