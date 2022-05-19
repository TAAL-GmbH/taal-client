package repository

import (
	"context"

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

func (r Repository) InsertKey(ctx context.Context, key server.Key) error {
	query := `INSERT INTO keys (api_key, private_key, public_key, address) VALUES ($1, $2, $3, $4);`
	_, err := r.db.ExecContext(ctx, query, key.ApiKey, key.PrivateKey, key.PublicKey, key.Address)

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

func (r Repository) InsertTransaction(ctx context.Context, txID string) error {

	query := `INSERT INTO transactions (id) VALUES ($1);`
	_, err := r.db.ExecContext(ctx, query, txID)

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
