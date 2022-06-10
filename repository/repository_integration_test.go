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

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	db, err = database.GetSQLiteDB("./localdata_test/db_test")
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
	t.Run("Insert key", func(t *testing.T) {
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
	})
}

func TestGetAllKeys(t *testing.T) {
	t.Run("Get all keys", func(t *testing.T) {
		is := is.New(t)
		err := prepareTestDatabase()
		is.NoErr(err)

		now := func() time.Time {
			return time.Date(2022, 5, 1, 10, 0, 0, 0, time.UTC)
		}

		repo := repository.NewRepository(db, now)
		ctx := context.Background()
		keys, err := repo.GetAllKeys(ctx)
		is.NoErr(err)

		expectedKeys := []server.Key{
			{
				ApiKey:     "11111",
				PublicKey:  "xskd023k3",
				PrivateKey: "2099n2dskd",
				Address:    "ke992kfj0",
				CreatedAt:  "2022-05-21 15:10:58.022+00:00",
			},
			{
				ApiKey:     "22222",
				PublicKey:  "adlkfsd9",
				PrivateKey: "xp3k0cj3m",
				Address:    "20fk2pdkf",
				CreatedAt:  "2022-05-24 15:10:58.022+00:00",
			},
		}

		is.Equal(expectedKeys, keys)
	})
}

func TestInsertTransaction(t *testing.T) {

	t.Run("Insert transaction", func(t *testing.T) {
		is := is.New(t)
		now := func() time.Time {
			return time.Date(2022, 6, 20, 10, 0, 0, 0, time.UTC)
		}

		repo := repository.NewRepository(db, now)
		ctx := context.Background()
		tx := server.Transaction{
			ID:        "7890abcde",
			ApiKey:    "api_key_3",
			DataBytes: 100,
			Filename:  "somepdf.pdf",
			Secret:    "secret",
		}
		err := repo.InsertTransaction(ctx, tx)
		is.NoErr(err)

		// time.Sleep(500)
		txsFromDB, err := repo.GetAllTransactions(ctx, 1)
		is.NoErr(err)

		is.Equal(1, len(txsFromDB))

		tx.CreatedAt = time.Date(2022, 6, 20, 10, 0, 0, 0, time.UTC).Format(repository.ISO8601)

		is.Equal(tx, txsFromDB[0])

	})

}

func TestGetTransactions(t *testing.T) {
	is := is.New(t)
	err := prepareTestDatabase()
	is.NoErr(err)

	now := func() time.Time {
		return time.Date(2022, 6, 1, 10, 0, 0, 0, time.UTC)
	}
	repo := repository.NewRepository(db, now)
	ctx := context.Background()

	tt := []struct {
		name        string
		hoursBack   int
		expectedTxs []server.Transaction
	}{
		{
			name:        "0 hours back",
			hoursBack:   0,
			expectedTxs: []server.Transaction{},
		},
		{
			name:      "30 days back",
			hoursBack: 24 * 30,
			expectedTxs: []server.Transaction{
				{
					ID:        "6A4410C3",
					ApiKey:    "api_key_2",
					DataBytes: 100,
					CreatedAt: "2022-05-25 15:10:58.022+00:00",
					Filename:  "somepicture2.png",
					Secret:    "secret",
				},
				{
					ID:        "2BDCFF23",
					ApiKey:    "api_key_1",
					DataBytes: 50,
					CreatedAt: "2022-05-23 15:10:58.022+00:00",
					Filename:  "textfile2.txt",
				},
				{
					ID:        "27EC83F0",
					ApiKey:    "api_key_1",
					DataBytes: 333,
					CreatedAt: "2022-05-12 22:10:58.022+00:00",
					Filename:  "somepicture5.png",
				},
				{
					ID:        "7650035F",
					ApiKey:    "api_key_2",
					DataBytes: 200,
					CreatedAt: "2022-05-12 15:10:58.022+00:00",
					Filename:  "somepicture1.png",
				},
				{
					ID:        "BA93B557",
					ApiKey:    "api_key_1",
					DataBytes: 100,
					CreatedAt: "2022-05-10 15:10:58.022+00:00",
					Filename:  "textfile1.txt",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			transactions, err := repo.GetAllTransactions(ctx, tc.hoursBack)
			is.NoErr(err)

			is.Equal(tc.expectedTxs, transactions)

		})
	}
}

func TestGetTransactionsStats(t *testing.T) {
	is := is.New(t)
	err := prepareTestDatabase()
	is.NoErr(err)

	to := time.Date(2022, 6, 1, 10, 0, 0, 0, time.UTC)

	repo := repository.NewRepository(db, nil)
	ctx := context.Background()

	tt := []struct {
		name        string
		from        time.Time
		expectedTxs []server.TransactionInfo
	}{
		{
			name:        "0 hours back",
			from:        to,
			expectedTxs: []server.TransactionInfo{},
		},
		{
			name: "30 days back",
			from: to.AddDate(0, 0, -30),
			expectedTxs: []server.TransactionInfo{
				{
					Timestamp: time.Date(2022, 5, 25, 0, 0, 0, 0, time.UTC),
					DataBytes: 100,
					Count:     1,
				},
				{
					Timestamp: time.Date(2022, 5, 23, 0, 0, 0, 0, time.UTC),
					DataBytes: 50,
					Count:     1,
				},
				{
					Timestamp: time.Date(2022, 5, 12, 0, 0, 0, 0, time.UTC),
					DataBytes: 533,
					Count:     2,
				},
				{
					Timestamp: time.Date(2022, 5, 10, 0, 0, 0, 0, time.UTC),
					DataBytes: 100,
					Count:     1,
				},
			},
		},
		{
			name: "30 days back, hour granularity",
			from: to.AddDate(0, 0, -30),
			expectedTxs: []server.TransactionInfo{
				{
					Timestamp: time.Date(2022, 5, 25, 0, 0, 0, 0, time.UTC),
					DataBytes: 100,
					Count:     1,
				},
				{
					Timestamp: time.Date(2022, 5, 23, 0, 0, 0, 0, time.UTC),
					DataBytes: 50,
					Count:     1,
				},
				{
					Timestamp: time.Date(2022, 5, 12, 0, 0, 0, 0, time.UTC),
					DataBytes: 533,
					Count:     2,
				},
				{
					Timestamp: time.Date(2022, 5, 10, 0, 0, 0, 0, time.UTC),
					DataBytes: 100,
					Count:     1,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			transactions, err := repo.GetTransactionInfo(ctx, tc.from, to, server.Day)
			is.NoErr(err)

			is.Equal(tc.expectedTxs, transactions)

		})
	}
}
