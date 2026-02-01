package repo

import (
	"context"
	"time"
	"url-shortener/database"
	shortener_models "url-shortener/shortener/models"
)

type Repo struct {
	dbManager database.DBManager
	txManager database.TXManager
}

func NewRepo(dbManager database.DBManager, txManager database.TXManager) *Repo {
	return &Repo{
		dbManager: dbManager,
		txManager: txManager,
	}
}

func (r *Repo) GetLinkByShort(ctx context.Context, short string, user_id int64) (string, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()

	query := "SELECT source FROM links WHERE short = $1 AND user_id = $2;"

	row := conn.QueryRow(ctx, query, short, user_id)

	var row_source string

	err = row.Scan(&row_source)
	if err != nil {
		return "", err
	}
	return row_source, nil
}

func (r *Repo) GetAllLinks(ctx context.Context, user_id int64) ([]shortener_models.LinkRecord, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query := "SELECT * FROM links WHERE user_id = $1;"

	rows, err := conn.Query(ctx, query, user_id)
	if err != nil {
		return nil, err
	}

	links := make([]shortener_models.LinkRecord, 0)

	for rows.Next() {
		var row_id int64
		var row_short string
		var row_source string
		var row_user_id int64
		var row_created_at time.Time
		var row_updated_at time.Time

		err = rows.Scan(&row_id, &row_short, &row_source, &row_user_id, &row_created_at, &row_updated_at)
		if err != nil {
			return nil, err
		}

		links = append(links, shortener_models.LinkRecord{
			ID:        row_id,
			Short:     row_short,
			Source:    row_source,
			UserID:    row_user_id,
			CreatedAt: row_created_at,
			UpdatedAt: row_updated_at,
		})
	}

	return links, nil
}

func (r *Repo) DeleteLink(ctx context.Context, short string, user_id int64) (bool, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	err = r.txManager.Execute(ctx, conn, func(ctx context.Context) error {
		query := "DELETE FROM links WHERE short = $1 AND user_id = $2;"

		_, err := conn.Exec(ctx, query, short, user_id)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repo) ShortLink(ctx context.Context, source string, short string, user_id int64) (bool, error) {
	conn, err := r.dbManager.GetConnection(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	err = r.txManager.Execute(ctx, conn, func(ctx context.Context) error {
		query := "INSERT INTO links(short, source, user_id) VALUES($1,$2,$3);"

		_, err := conn.Exec(ctx, query, short, source, user_id)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
