package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimk00z/go-musthave-diploma/config"
	api "github.com/dimk00z/go-musthave-diploma/internal/controller/http/api"
	"github.com/dimk00z/go-musthave-diploma/pkg/httpserver"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
	"github.com/dimk00z/go-musthave-diploma/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// HTTP Server
	handler := gin.New()
	// api.NewRouter(handler, l, translationUseCase)
	api.NewRouter(handler, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

	}
	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
