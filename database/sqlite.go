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
	dir, _ := filepath.Split(dbFilePath)

	if err := os.Mkdir(dir, 0755); err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	sqliteDb, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	dbPathToRemove = dir

	return sqlx.NewDb(sqliteDb, "sqlite"), nil
}

func RemoveSQLiteDB() error {
	err := os.RemoveAll(dbPathToRemove)
	if err != nil {
		return errors.Wrap(err, "unable to remove sqlite db")
	}

	return nil
}
