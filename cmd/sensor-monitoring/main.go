package main

import (
	"log"
	sensor_monitoring "sensor-monitoring/internal/app/sensor-monitoring"
	"sensor-monitoring/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	app := sensor_monitoring.New(cfg)

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
}
