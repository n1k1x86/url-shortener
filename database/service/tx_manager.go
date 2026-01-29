package service

import (
	"context"
	"database/sql"
)

type TXManager struct {
	db *sql.DB
}

func (t *TXManager) Execute(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
