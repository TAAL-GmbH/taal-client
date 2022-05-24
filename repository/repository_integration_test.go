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

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	"github.com/pkg/errors"
)

var (
	db       *sqlx.DB
	fixtures *testfixtures.Loader
)

const (
	dbFolder         = "localdata_test"
	sqLiteDBFilename = "db_test"
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

	fixtures, err = testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect("sqlite"),              // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata/fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		return -1, errors.Wrap(err, "failed to run fixtures")
	}
	return m.Run(), nil
}

func prepareTestDatabase() error {
	err := fixtures.Load()
	if err != nil {
		return errors.Wrap(err, "loading fixtures failed")
	}

	return nil
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

func TestGetTransactions(t *testing.T) {
	is := is.New(t)
	err := prepareTestDatabase()
	is.NoErr(err)

	now := func() time.Time {
		return time.Date(2022, 5, 1, 10, 0, 0, 0, time.UTC)
	}
	repo := repository.NewRepository(db, now)
	ctx := context.Background()

	transactions, err := repo.GetAllTransactions(ctx)
	is.NoErr(err)

	expectedTransactions := []server.Transaction{
		{
			ID:        "12345abcde",
			ApiKey:    "api_key_1",
			DataBytes: 100,
			CreatedAt: "2022-05-23 15:10:58.022+00:00",
		},
	}

	is.Equal(expectedTransactions, transactions)

}
