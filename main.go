package main

import (
	"auth-service/config"
	"auth-service/internal/app"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatalf("error loading config error:%s", configErr)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Interrupt)
	defer cancel()

	server := app.New()
	err := server.Start(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
