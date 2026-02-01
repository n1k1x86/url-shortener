package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	auth_repo "url-shortener/auth/repo"
	auth_service "url-shortener/auth/service"
	"url-shortener/config"
	db_service "url-shortener/database/service"
	"url-shortener/server/service"
	shortener_repo "url-shortener/shortener/repo"
	shortener_service "url-shortener/shortener/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbManager := db_service.NewDBManager(&cfg.Database)
	err = dbManager.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	txManager := db_service.NewTXManager()

	authrepo := auth_repo.NewRepo(dbManager, txManager)
	authService := auth_service.NewService(&cfg.Auth, authrepo)

	access, refresh, _, _ := authService.GenerateTokenPair(ctx, 1, "user")
	fmt.Println("access: ", access)
	fmt.Println("refresh: ", refresh)

	shortenerRepo := shortener_repo.NewRepo(dbManager, txManager)
	shortenerService := shortener_service.NewService(shortenerRepo)

	serv := service.NewHTTPServer(ctx, authService, shortenerService, ":8080")
	serv.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()
	serv.Stop(timeoutCtx)
	dbManager.Close()

	log.Println("application was gracefully shutted down")
}
