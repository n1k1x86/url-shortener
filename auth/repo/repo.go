package repo

import (
	"context"
	"database/sql"
	"time"
	"url-shortener/database"
)

type Repository struct {
	dbManager database.DBManager
	txManager database.TXManager
}

func NewRepo(dbManager database.DBManager, txManager database.TXManager) *Repository {
	return &Repository{dbManager: dbManager, txManager: txManager}
}

func (r *Repository) InsertRefreshToken(ctx context.Context, refreshHash, jti string, userID int64) error {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = r.txManager.Execute(ctx, conn, func(ctx context.Context) error {
		query := "INSERT INTO refresh_tokens(jti, hash, user_id, replaced_by, is_revoked) VALUES($1, $2, $3, $4, $5)"
		_, err := conn.Exec(ctx, query, jti, refreshHash, userID, "", false)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RevokeToken(ctx context.Context, jtiNew, jtiOld string) error {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = r.txManager.Execute(ctx, conn, func(ctx context.Context) error {
		query := "UPDATE refresh_tokens SET is_revoked = $1, updated_at = $2, replaced_by = $3 WHERE jti = $4"
		_, err := conn.Exec(ctx, query, true, time.Now(), jtiNew, jtiOld)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) IsTokenRevoked(ctx context.Context, jti string) (bool, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	query := "SELECT is_revoked FROM refresh_tokens WHERE jti = $1;"
	row := conn.QueryRow(ctx, query, jti)

	var isRevoked bool
	err = row.Scan(&isRevoked)
	if err != nil {
		return false, err
	}

	return isRevoked, nil
}

func (r *Repository) IsUserExist(ctx context.Context, login string, userID int64) (bool, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	query := "SELECT login FROM users WHERE id = $1 and login = $2;"
	row := conn.QueryRow(ctx, query, userID, login)

	var loginNull sql.NullString
	err = row.Scan(&loginNull)
	if err != nil {
		return false, err
	}

	return loginNull.Valid, nil
}
