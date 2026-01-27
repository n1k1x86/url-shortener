package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func NewRouter() *chi.Mux {
	m := chi.NewRouter()
	m.Get("/hello", HelloHandler)
	return m
}
