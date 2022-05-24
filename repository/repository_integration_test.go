package repository_test

import (
	"context"

	"fmt"
	"os"
	"taal-client/database"
	"taal-client/repository"
	"taal-client/server"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	"github.com/pkg/errors"
)

var (
	db *sqlx.DB
)

const (
	dbFolder         = "localdata_test"
	sqLiteDBFilename = "db"
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	db, err = database.GetSQLiteDB(dbFolder, sqLiteDBFilename)
	if err != nil {
		return -1, errors.Wrap(err, "failed to set up db")
	}

	err = database.RunMigrationsSQLite(db)
	if err != nil {
		return -1, errors.Wrap(err, "failed to run migrations")
	}

	defer func() {
		db.Close()
		err = database.RemoveSQLiteDB()
	}()

	return m.Run(), nil
}

func TestInsertKey(t *testing.T) {
	is := is.New(t)

	now := func() time.Time {
		return time.Date(2022, 5, 1, 10, 0, 0, 0, time.UTC)
	}

	repo := repository.NewRepository(db, now)
	ctx := context.Background()
	key := server.Key{
		ApiKey:     "test_api_key",
		PublicKey:  "test_public_key",
		PrivateKey: "test_private_key",
		Address:    "test_address",
		RevokedAt:  nil,
	}
	err := repo.InsertKey(ctx, key)
	is.NoErr(err)

	keyFromDB, err := repo.GetKey(ctx, "test_api_key")
	is.NoErr(err)

	key.CreatedAt = time.Date(2022, 5, 1, 10, 0, 0, 0, time.UTC).Format(repository.ISO8601)

	is.Equal(key, keyFromDB)
}
