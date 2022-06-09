package repository

type TransactionInfo struct {
	Timestamp string `db:"timestamp" json:"timestamp"`
	Count     int    `db:"count" json:"count"`
	DataBytes int    `db:"data_bytes" json:"data_bytes"`
}
