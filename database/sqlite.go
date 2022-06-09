package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	dbPathToRemove string
)

func GetSQLiteDB(dbFilePath string) (*sqlx.DB, error) {
	if err := os.Mkdir(filepath.Dir(dbFilePath), 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	sqliteDb, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	dbPathToRemove = dbFilePath

	return sqlx.NewDb(sqliteDb, "sqlite3"), nil
}

func RemoveSQLiteDB() error {
	err := os.RemoveAll(dbPathToRemove)
	if err != nil {
		return errors.Wrap(err, "unable to remove sqlite db")
	}

	return nil
}
