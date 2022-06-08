package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"taal-client/server"
)

type Repository struct {
	db  *sqlx.DB
	now func() time.Time
}

func NewRepository(db *sqlx.DB, now func() time.Time) Repository {
	return Repository{
		db:  db,
		now: now,
	}
}

const ISO8601 = "2006-01-02T15:04:05.999Z"

func (r Repository) InsertKey(ctx context.Context, key server.Key) error {
	createdAt := r.now().UTC().Format(ISO8601)

	query := `INSERT INTO keys (created_at, api_key, private_key, public_key, address) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.db.ExecContext(ctx, query, createdAt, key.ApiKey, key.PrivateKey, key.PublicKey, key.Address)

	return err
}

func (r Repository) GetKey(ctx context.Context, apiKey string) (server.Key, error) {
	query := `SELECT * FROM keys WHERE api_key = $1 LIMIT 1;`

	key := server.Key{}

	err := r.db.GetContext(ctx, &key, query, apiKey)
	if err != nil {
		return server.Key{}, err
	}

	return key, nil
}

func (r Repository) GetAllKeys(ctx context.Context) ([]server.Key, error) {
	query := `SELECT * FROM keys WHERE revoked_at IS NULL;`

	keys := make([]server.Key, 0)

	err := r.db.SelectContext(ctx, &keys, query)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r Repository) InsertTransaction(ctx context.Context, tx server.Transaction) error {
	createdAt := r.now().UTC().Format(ISO8601)
	query := `INSERT INTO transactions (created_at, id, api_key, data_bytes, filename) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.db.ExecContext(ctx, query, createdAt, tx.ID, tx.ApiKey, tx.DataBytes, tx.Filename)

	return err
}

func (r Repository) GetAllTransactions(ctx context.Context, hoursBack int) ([]server.Transaction, error) {
	now := r.now()

	timeBack := now.Add(-1 * time.Duration(hoursBack) * time.Hour).UTC().Format(ISO8601)
	query := `SELECT * FROM transactions WHERE created_at >= $1 ORDER BY created_at DESC;`

	txs := make([]server.Transaction, 0)

	err := r.db.SelectContext(ctx, &txs, query, timeBack)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (r Repository) Health(ctx context.Context) error {
	return r.db.Ping()
}
