package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	auth_service "url-shortener/auth/service"
	"url-shortener/config"
	"url-shortener/server/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	authService := auth_service.NewService(&cfg.Auth)
	access, refresh, err := authService.GenerateTokenPair(123123, "my-test-login")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("access: %s\n", access)
	fmt.Printf("refresh: %s\n", refresh)

	serv := service.NewHTTPServer(":8080")
	serv.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()
	serv.Stop(timeoutCtx)
	log.Println("application was gracefully shutted down")
}
