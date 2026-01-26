package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/server/service"
)

func main() {
	serv := service.NewHTTPServer(":8082")
	serv.Run()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeoutCancel()
	serv.Stop(timeoutCtx)
	log.Println("application was gracefully shutted down")
}
