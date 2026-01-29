package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager interface {
	Connect(ctx context.Context) (*pgxpool.Pool, error)
	GetConnection(ctx context.Context) (*pgxpool.Conn, error)
	Close()
}

type TXManager interface {
	Execute(ctx context.Context, fn func(tx *sql.Tx)) error
}
