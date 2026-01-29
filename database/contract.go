package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager interface {
	Connect() (*pgxpool.Pool, error)
	GetConnection() (*pgxpool.Conn, error)
	Close() error
}

type TXManager interface {
	Execute(ctx context.Context, fn func(tx *sql.Tx)) error
}
