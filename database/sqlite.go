package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"taal-client/settings"
)

var (
	dbPathToRemove string
)

func GetSQLiteDB(dbPath string, dbFilename string) (*sqlx.DB, error) {
	if err := os.Mkdir(dbPath, 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	dbFilePath := settings.Get("dbFilename")

	sqliteDb, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	dbPathToRemove = dbPath

	return sqlx.NewDb(sqliteDb, "sqlite3"), nil
}

func RemoveSQLiteDB() error {
	err := os.RemoveAll(fmt.Sprintf("./%s/", dbPathToRemove))

	if err != nil {
		return errors.Wrap(err, "unable to remove sqlite db")
	}

	return nil
}
