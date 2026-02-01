package router

import (
	"context"
	"net/http"
	"url-shortener/auth"
	"url-shortener/shortener"

	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func NewRouter(ctx context.Context, authService auth.Auth, shortener shortener.Shortener) *chi.Mux {
	m := chi.NewRouter()
	m.Get("/hello", HelloHandler)
	m.Group(func(r chi.Router) {
		r.Use(AuthMiddleware(ctx, authService))
		r.Get("/links/get/{short}", GetLinkByShort(ctx, shortener))
		r.Get("/links/get", GetAllLinks(ctx, shortener))
		r.Delete("/links/delete/{short}", DeleteLink(ctx, shortener))
		r.Post("/links/short", ShortLink(ctx, shortener))
	})
	return m
}
