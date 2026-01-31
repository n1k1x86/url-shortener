package shortener

import (
	"context"
	shortener_models "url-shortener/shortener/models"
)

type Shortener interface {
	GetLinkByShort(ctx context.Context, short string, user_id int64) (string, error)
	GetAllLinks(ctx context.Context, user_id int64) ([]shortener_models.LinkRecord, error)
	DeleteLink(ctx context.Context, short string, user_id int64) (bool, error)
	ShortLink(ctx context.Context, source string, short string, user_id int64) (bool, error)
}
