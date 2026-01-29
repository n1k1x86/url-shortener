package service

import (
	"context"
	"fmt"
	db_config "url-shortener/database/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager struct {
	pool *pgxpool.Pool
	cfg  *db_config.Config
}

func (d *DBManager) buildDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", d.cfg.DBConfig.Username, d.cfg.DBConfig.Password, d.cfg.DBConfig.Server, d.cfg.DBConfig.Database)
}

func (d *DBManager) Connect(ctx context.Context) error {

	poolCfg, err := pgxpool.ParseConfig(d.buildDSN())
	if err != nil {
		return err
	}

	poolCfg.MaxConns = d.cfg.ConnConfig.MaxConns
	poolCfg.MinConns = d.cfg.ConnConfig.MinConns
	poolCfg.MaxConnLifetime = d.cfg.ConnConfig.MaxConnLifeTime
	poolCfg.MaxConnIdleTime = d.cfg.ConnConfig.MaxConnIdleTime
	poolCfg.HealthCheckPeriod = d.cfg.ConnConfig.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return err
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, d.cfg.PoolConfig.PingTimeout)
	defer cancel()

	if err = pool.Ping(ctxTimeout); err != nil {
		pool.Close()
		return err
	}

	d.pool = pool
	return nil
}

func (d *DBManager) GetConnection(ctx context.Context) (*pgxpool.Conn, error) {
	if d.pool == nil {
		return nil, fmt.Errorf("pool is nil")
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, d.cfg.PoolConfig.AcquireTimeout)
	defer cancel()

	conn, err := d.pool.Acquire(ctxTimeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (d *DBManager) Close() {
	if d.pool != nil {
		d.pool.Close()
		d.pool = nil
	}
}
