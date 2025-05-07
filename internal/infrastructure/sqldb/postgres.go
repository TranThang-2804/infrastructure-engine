package sqldb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// PostgresInfraImpl is the concrete implementation of the PostgresInfra interface.
type PostgresDb struct {
	db *sql.DB
}

// NewPostgresInfra creates a new instance of PostgresInfraImpl.
func NewPostgresInfra(dsn string) (SqlDb, error) {
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresDb{db: db}, nil
}

// Exec executes a query that does not return rows (e.g., INSERT, UPDATE, DELETE).
func (p *PostgresDb) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows (e.g., SELECT).
func (p *PostgresDb) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that returns a single row.
func (p *PostgresDb) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

// BeginTx starts a new database transaction.
func (p *PostgresDb) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return p.db.BeginTx(ctx, opts)
}

// Close closes the database connection.
func (p *PostgresDb) Close() error {
	return p.db.Close()
}

// Ping verifies that the database connection is alive.
func (p *PostgresDb) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}
