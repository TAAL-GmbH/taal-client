package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"taal-client/settings"
)

var (
	dbPathToRemove string
)

func GetSQLiteDB() (*sqlx.DB, error) {
	dbFilePath := settings.Get("dbFilename")

	dir, _ := filepath.Split(dbFilePath)

	if err := os.Mkdir(dir, 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	sqliteDb, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	dbPathToRemove = dir

	return sqlx.NewDb(sqliteDb, "sqlite3"), nil
}

func RemoveSQLiteDB() error {
	err := os.RemoveAll(fmt.Sprintf("./%s/", dbPathToRemove))

	if err != nil {
		return errors.Wrap(err, "unable to remove sqlite db")
	}

	return nil
}
