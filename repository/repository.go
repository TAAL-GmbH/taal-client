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
	query := `SELECT * FROM keys WHERE api_key = $1 LIMIT 1`

	key := server.Key{}

	err := r.db.GetContext(ctx, &key, query, apiKey)
	if err != nil {
		return server.Key{}, err
	}

	return key, nil
}
