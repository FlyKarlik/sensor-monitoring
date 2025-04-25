package main

import (
	"log"
	"sensor-monitoring/internal/app"
	"sensor-monitoring/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	app := app.New(cfg)

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
}
