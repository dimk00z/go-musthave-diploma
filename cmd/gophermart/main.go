package main

import (
	"log"
	"os"

	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/app"
	"github.com/rs/zerolog"
)

func main() {
	// Configuration
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)

	logger.Log().Msg("starting App")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
