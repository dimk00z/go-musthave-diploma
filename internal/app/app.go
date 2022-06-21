package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimk00z/go-musthave-diploma/config"
	api "github.com/dimk00z/go-musthave-diploma/internal/controller/http/api"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase/repo"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase/webapi"
	"github.com/dimk00z/go-musthave-diploma/internal/worker"
	"github.com/dimk00z/go-musthave-diploma/pkg/httpserver"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/dimk00z/go-musthave-diploma/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New("debug")
	l.Debug(cfg)
	log.Println(cfg)
	// Migrate
	doMigrations(cfg.PG.URL, l)
	// Workers pool
	wp := worker.GetWorkersPool(cfg.Workers, l)
	defer wp.Close()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		wp.Run(ctx)
	}()

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	l.Debug("app - Connected to DB")
	defer pg.Close()

	// Use case
	gophermartUseCase := usecase.New(
		repo.New(pg),
		webapi.New(cfg),
		l,
	)

	// HTTP Server
	mainRouter := gin.Default()
	api.NewRouter(mainRouter, l, gophermartUseCase, cfg)
	httpServer := httpserver.New(mainRouter, httpserver.Addr(cfg.HTTP.RunAddress))
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
	err = httpServer.Shutdown(ctx, cancel)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	// defer cancel()

}
