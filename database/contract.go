package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager interface {
	Connect(ctx context.Context) error
	GetConnection(ctx context.Context) (*pgxpool.Conn, error)
	Close()
}

type TXManager interface {
	Execute(ctx context.Context, conn *pgxpool.Conn, fn func(ctx context.Context) error) error
}
