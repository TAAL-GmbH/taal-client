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
	db sqlx.DB
}

func NewRepository(db sqlx.DB) Repository {
	return Repository{db: db}
}

const ISO8601 = "2006-01-02T15:04:05.999Z"

func (r Repository) InsertKey(ctx context.Context, key server.Key) error {
	createdAt := time.Now().UTC().Format(ISO8601)

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

	query := `INSERT INTO transactions (id, api_key) VALUES ($1, $2);`
	_, err := r.db.ExecContext(ctx, query, tx.ID, tx.ApiKey)

	return err
}

func (r Repository) GetAllTransactions(ctx context.Context) ([]server.Transaction, error) {
	query := `SELECT * FROM transactions;`

	txs := make([]server.Transaction, 0)

	err := r.db.SelectContext(ctx, &txs, query)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
