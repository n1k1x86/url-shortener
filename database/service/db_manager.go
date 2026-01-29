package service

import (
	db_config "url-shortener/database/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager struct {
	pool *pgxpool.Pool
	cfg  *db_config.Config
}

func (d *DBManager) Connect() (*pgxpool.Pool, error) {
	return nil, nil
}

func (d *DBManager) GetConnection() (*pgxpool.Conn, error) {
	return nil, nil
}

func (d *DBManager) Close() error {
	return nil
}
