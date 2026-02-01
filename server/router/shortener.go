package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"url-shortener/server/models"
	"url-shortener/shortener"

	"github.com/go-chi/chi/v5"
)

func GetLinkByShort(ctx context.Context, shortener shortener.Shortener) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int64)
		short := chi.URLParam(r, "short")

		short, err := shortener.GetLinkByShort(ctx, short, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(short))
	})
}

func GetAllLinks(ctx context.Context, shortener shortener.Shortener) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int64)

		links, err := shortener.GetAllLinks(ctx, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		var resp models.GetAllLinksResponse = models.GetAllLinksResponse{
			Data: links,
		}

		data, err := json.Marshal(&resp)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	})
}

func DeleteLink(ctx context.Context, shortener shortener.Shortener) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int64)

		short := chi.URLParam(r, "short")

		_, err := shortener.DeleteLink(ctx, short, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func ShortLink(ctx context.Context, shortener shortener.Shortener) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int64)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		var req models.ShortLinkBody
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		_, err = shortener.ShortLink(ctx, req.Source, req.Short, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
