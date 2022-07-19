package repository_test

import (
	"context"
	"log"
	"strconv"

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

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/pkg/errors"
)

var (
	db       *sqlx.DB
	fixtures *testfixtures.Loader

	dbuser     = "postgres"
	dbpassword = "secret"
	dbname     = "postgres_test"
	dbport     = 5433
	dbdialect  = "postgres"
)

func TestMain(m *testing.M) {
	var code int
	var err error

	db := os.Getenv("DB")

	switch db {
	case "POSTGRES":
		code, err = runPostgres(m)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(code)
	case "SQLITE":
		code, err = runSqLite(m)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(code)
	default:
		log.Println("no db set with env var 'DB' - skipping integration tests")
	}
}

func runSqLite(m *testing.M) (code int, err error) {
	db, err = database.GetSQLiteDB("./localdata/testdb")
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

	fixtures.Load()
	if err != nil {
		return -1, errors.Wrap(err, "failed to load fixtures")
	}

	return m.Run(), nil
}

func runPostgres(m *testing.M) (code int, err error) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=" + dbuser,
			"POSTGRES_PASSWORD=" + dbpassword,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: strconv.Itoa(dbport)},
			},
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	defer pool.Purge(resource)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {

		db, err = database.GetPostgreSqlDB("localhost", dbport, dbuser, dbpassword, dbname)

		if err != nil {
			return err
		}

		err = db.Ping()
		if err != nil {
			return err
		}

		return nil

	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	err = database.RunMigrationsPostgreSQL(db)
	if err != nil {
		return -1, errors.Wrap(err, "failed to run migrations")
	}

	defer db.Close()

	fixtures, err = testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect(dbdialect),             // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata/fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		return -1, errors.Wrap(err, "failed to run fixtures")
	}

	fixtures.Load()
	if err != nil {
		return -1, errors.Wrap(err, "failed to load fixtures")
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
				ApiKey:     "api_key_1",
				PublicKey:  "xskd023k3",
				PrivateKey: "2099n2dskd",
				Address:    "ke992kfj0",
				CreatedAt:  "2022-05-21 15:10:58.022Z",
			},
			{
				ApiKey:     "api_key_2",
				PublicKey:  "adlkfsd9",
				PrivateKey: "xp3k0cj3m",
				Address:    "20fk2pdkf",
				CreatedAt:  "2022-05-24 15:10:58.022Z",
			},
			{
				ApiKey:     "api_key_4",
				PublicKey:  "7a2f1cb9",
				PrivateKey: "cb7168ab",
				Address:    "5ec39af2",
				CreatedAt:  "2022-06-10 15:10:58.022Z",
			},
		}

		is.Equal(expectedKeys, keys)
	})
}

func TestGetAllKeyUsages(t *testing.T) {
	t.Run("Get all keys Usage", func(t *testing.T) {
		is := is.New(t)
		err := prepareTestDatabase()
		is.NoErr(err)

		now := func() time.Time {
			return time.Date(2022, 5, 1, 10, 0, 0, 0, time.UTC)
		}

		repo := repository.NewRepository(db, now)
		ctx := context.Background()
		keys, err := repo.GetAllKeysUsage(ctx)
		is.NoErr(err)

		expectedKeys := []server.KeyUsage{
			{
				Key: server.Key{
					ApiKey:     "api_key_1",
					PublicKey:  "xskd023k3",
					PrivateKey: "2099n2dskd",
					Address:    "ke992kfj0",
					CreatedAt:  "2022-05-21 15:10:58.022Z",
				},
				DataBytes: 523,
			},
			{
				Key: server.Key{
					ApiKey:     "api_key_2",
					PublicKey:  "adlkfsd9",
					PrivateKey: "xp3k0cj3m",
					Address:    "20fk2pdkf",
					CreatedAt:  "2022-05-24 15:10:58.022Z",
				},
				DataBytes: 300,
			},
			{
				Key: server.Key{
					ApiKey:     "api_key_4",
					PublicKey:  "7a2f1cb9",
					PrivateKey: "cb7168ab",
					Address:    "5ec39af2",
					CreatedAt:  "2022-06-10 15:10:58.022Z",
				},
				DataBytes: 0,
			},
		}

		is.Equal(expectedKeys, keys)
	})
}

func TestDeactivateKey(t *testing.T) {
	t.Run("Deactivate key", func(t *testing.T) {
		is := is.New(t)
		now := func() time.Time {
			return time.Date(2022, 6, 20, 10, 0, 0, 0, time.UTC)
		}

		repo := repository.NewRepository(db, now)
		ctx := context.Background()

		err := repo.DeactivateKey(ctx, "api_key_2")
		is.NoErr(err)

		key, err := repo.GetKey(ctx, "api_key_2")
		is.NoErr(err)

		is.Equal("2022-06-20T10:00:00Z", *key.RevokedAt)

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
			IsHash:    true,
		}
		err := repo.InsertTransaction(ctx, tx)
		is.NoErr(err)

		txsFromDB, err := repo.GetAllTransactions(ctx, false, 1)
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
		all         bool
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
					CreatedAt: "2022-05-25 15:10:58.022Z",
					Filename:  "somepicture2.png",
				},
				{
					ID:        "2BDCFF23",
					ApiKey:    "api_key_1",
					DataBytes: 50,
					CreatedAt: "2022-05-23 15:10:58.022Z",
					Filename:  "textfile2.txt",
					Secret:    "1234",
				},
				{
					ID:        "27EC83F0",
					ApiKey:    "api_key_1",
					DataBytes: 333,
					CreatedAt: "2022-05-12 22:10:58.022Z",
					Filename:  "somepicture5.png",
				},
				{
					ID:        "7650035F",
					ApiKey:    "api_key_2",
					DataBytes: 200,
					CreatedAt: "2022-05-12 15:10:58.022Z",
					Filename:  "somepicture1.png",
				},
				{
					ID:        "BA93B557",
					ApiKey:    "api_key_1",
					DataBytes: 100,
					CreatedAt: "2022-05-10 15:10:58.022Z",
					Filename:  "textfile1.txt",
				},
			},
		},
		{
			name: "all",
			all:  true,
			expectedTxs: []server.Transaction{
				{
					ID:        "6A4410C3",
					ApiKey:    "api_key_2",
					DataBytes: 100,
					CreatedAt: "2022-05-25 15:10:58.022Z",
					Filename:  "somepicture2.png",
				},
				{
					ID:        "2BDCFF23",
					ApiKey:    "api_key_1",
					DataBytes: 50,
					CreatedAt: "2022-05-23 15:10:58.022Z",
					Filename:  "textfile2.txt",
					Secret:    "1234",
				},
				{
					ID:        "27EC83F0",
					ApiKey:    "api_key_1",
					DataBytes: 333,
					CreatedAt: "2022-05-12 22:10:58.022Z",
					Filename:  "somepicture5.png",
				},
				{
					ID:        "7650035F",
					ApiKey:    "api_key_2",
					DataBytes: 200,
					CreatedAt: "2022-05-12 15:10:58.022Z",
					Filename:  "somepicture1.png",
				},
				{
					ID:        "BA93B557",
					ApiKey:    "api_key_1",
					DataBytes: 100,
					CreatedAt: "2022-05-10 15:10:58.022Z",
					Filename:  "textfile1.txt",
				},
				{
					ID:        "2C34AE2C",
					ApiKey:    "api_key_1",
					DataBytes: 40,
					CreatedAt: "2022-04-28 15:10:58.022Z",
					Filename:  "textfile0.txt",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			transactions, err := repo.GetAllTransactions(ctx, tc.all, tc.hoursBack)
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
