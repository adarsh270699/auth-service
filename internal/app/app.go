package app

import (
	"auth-service/config"
	"auth-service/internal/api/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: routers.LoadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    ":" + config.Config.ServerConfig.Port,
		Handler: a.router,
	}

	ch := make(chan error, 1)
	go func() {
		log.Println("server running on port:", config.Config.ServerConfig.Port)
		serverErr := server.ListenAndServe()
		if serverErr != nil {
			ch <- fmt.Errorf("error starting the server, error:%w", serverErr)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		fmt.Println("server shutting down")
		return server.Shutdown(timeout)
	}
}
