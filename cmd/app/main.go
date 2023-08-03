package main

import (
	"go-mongo-crud-rest-api/internal/app"
	"go-mongo-crud-rest-api/internal/config"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
