package main

import (
	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/app"
	"github.com/dimk00z/go-musthave-diploma/pkg/logger"
)

func main() {
	// Configuration
	cfg := config.NewConfig()
	l := logger.New("DEBUG")
	l.Debug(cfg)
	// Run
	app.Run(cfg, l)
}
