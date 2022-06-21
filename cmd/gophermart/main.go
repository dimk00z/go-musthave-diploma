package main

import (
	"log"

	"github.com/dimk00z/go-musthave-diploma/config"
	"github.com/dimk00z/go-musthave-diploma/internal/app"
)

func main() {
	log.Println("log test")
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	// Run
	app.Run(cfg)
}
