package service

import (
	"context"
	"errors"
	"log"
	"net/http"
	"url-shortener/auth"
	"url-shortener/server/router"
	"url-shortener/shortener"
)

type Service struct {
	server *http.Server
}

func (s *Service) Run() {
	go s.run()
}

func (s *Service) Stop(ctx context.Context) {
	s.server.Shutdown(ctx)
}

func (s *Service) run() {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("ERROR: panic was recovered %s", r)
		}
	}()
	log.Println("server was started")
	err := s.server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("server was stopped")
			return
		}
		log.Printf("ERROR: %s", err.Error())
	}
}

func NewHTTPServer(ctx context.Context, authService auth.Auth, shortener shortener.Shortener, addr string) *Service {
	r := router.NewRouter(ctx, authService, shortener)

	return &Service{
		server: &http.Server{Addr: addr, Handler: r},
	}
}
