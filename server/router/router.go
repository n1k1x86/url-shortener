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
	return m
}
