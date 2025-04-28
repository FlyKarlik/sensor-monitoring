package main

import (
	"log"
	"os"
	"sensor-monitoring/internal/app/migrator"
	"sensor-monitoring/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("Usage: %s up|down", os.Args[0])
	}

	app, err := migrator.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}

	switch args[0] {
	case "up":
		err = app.Up()
	case "down":
		err = app.Down()
	default:
		log.Fatalf("Usage: %s up|down", os.Args[0])
	}

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
