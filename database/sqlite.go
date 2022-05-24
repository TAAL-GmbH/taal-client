package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

	sqliteDb, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s", dbPath, dbFilename))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB")
	}

	dbPathToRemove = dbPath

	return sqlx.NewDb(sqliteDb, "sqlite3"), nil
}
