package sqldb

import (
	"context"
	"database/sql"
)

// PostgresInfra defines the interface for interacting with a PostgreSQL database.
type SqlDb interface {
	// Execute a query that does not return rows (e.g., INSERT, UPDATE, DELETE).
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Execute a query that returns rows (e.g., SELECT).
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// Execute a query that returns a single row.
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row

	// Begin a new database transaction.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Close the database connection.
	Close() error

	// Ping the database to ensure the connection is alive.
	Ping(ctx context.Context) error
}
