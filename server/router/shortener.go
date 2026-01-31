package router

import (
	"context"
	"fmt"
	"net/http"
	"url-shortener/shortener"
)

func GetLinkByShort(ctx context.Context, shortener shortener.Shortener) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int64)
		short := r.URL.Query().Get("short")

		short, err := shortener.GetLinkByShort(ctx, short, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(short))
	})
}

// GetLinkByShort(ctx context.Context, short string, user_id int64) (string, error)
// GetAllLinks(ctx context.Context, user_id int64) ([]shortener_models.LinkRecord, error)
// DeleteLink(ctx context.Context, short string, user_id int64) (bool, error)
// ShortLink(ctx context.Context, source string, short string, user_id int64) (bool, error)
