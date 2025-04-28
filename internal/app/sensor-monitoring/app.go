package sensor_monitoring

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sensor-monitoring/internal/config"
	"sensor-monitoring/internal/delivery/sync/http/handler"
	httpserver "sensor-monitoring/internal/delivery/sync/http/server"
	"sensor-monitoring/internal/repository"
	"sensor-monitoring/internal/usecase"
	"sensor-monitoring/pkg/database"
	"sensor-monitoring/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logger.New(a.cfg.SensorServiceConfig.LogLevel)
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create logger", err)
		return err
	}

	logger.Info("App", "App.Start", "Starting sensor service...")

	postgresDB, err := database.NewPostgresDB(&a.cfg.Infrastructure.Postgres)
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create postgres db", err)
		return err
	}
	repo := repository.New(postgresDB, logger)
	usecase := usecase.New(repo, logger)
	handler := handler.New(logger, usecase)
	httpServer := httpserver.NewHTTPServer(logger, a.cfg, handler)

	go func() {
		logger.Info("App", "App.Start", "Starting http server...", a.cfg.ServiceHost, a.cfg.ServicePort)
		if err := httpServer.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("App", "App.Start", "Failed to start http server", err)
			os.Exit(1)
		}
	}()

	signalHandler(ctx, logger)
	shutdown(context.WithoutCancel(ctx), logger, postgresDB, httpServer)
	return nil
}

func signalHandler(ctx context.Context, logger logger.Logger) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		close(signalChan)
	}()

	select {
	case <-ctx.Done():
		logger.Info("App", "App.signalHandler", "Context cancelled")
		return
	case <-signalChan:
		logger.Info("App", "App.signalHandler", "Signal received")
		return
	}
}

func shutdown(
	ctx context.Context,
	logger logger.Logger,
	postgresDB *pgxpool.Pool,
	httpServer *httpserver.HTTPServer,
) error {
	postgresDB.Close()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("App", "App.shutdown", "Failed to shutdown http server", err)
		return err
	}

	return nil
}
