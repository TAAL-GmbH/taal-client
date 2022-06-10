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

func (r Repository) GetTransactionInfo(ctx context.Context, from time.Time, to time.Time, granularity server.Granularity) ([]server.TransactionInfo, error) {

	query := `SELECT SUBSTR(created_at, 0, $1) AS timestamp, count(*) as count, sum(data_bytes) AS data_bytes FROM transactions WHERE created_at > $2 AND created_at < $3 GROUP BY timestamp ORDER BY timestamp DESC;`

	txs := make([]TransactionInfo, 0)
	position, format := granularitySecondsToPositionAndFormat(granularity)
	err := r.db.SelectContext(ctx, &txs, query, position, from.Format(ISO8601), to.Format(ISO8601))
	if err != nil {
		return nil, err
	}

	txInfos := make([]server.TransactionInfo, len(txs))

	for i, tx := range txs {
		timestamp, err := time.Parse(format, tx.Timestamp)
		if err != nil {
			return nil, err
		}
		txInfos[i] = server.TransactionInfo{
			Timestamp: timestamp,
			Count:     tx.Count,
			DataBytes: tx.DataBytes,
		}
	}

	return txInfos, nil
}

func (r Repository) Health(ctx context.Context) error {
	return r.db.Ping()
}

func granularitySecondsToPositionAndFormat(granularitySeconds server.Granularity) (int, string) {
	switch granularitySeconds {
	case server.None:
		return 20, "2006-01-02T15:04:05"
	case server.Minute:
		return 17, "2006-01-02T15:04"
	case server.Hour:
		return 14, "2006-01-02T15"
	}

	// Day
	return 11, "2006-01-02"
}
