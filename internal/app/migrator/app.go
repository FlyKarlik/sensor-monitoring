package migrator

import (
	"database/sql"
	"sensor-monitoring/internal/config"
	"sensor-monitoring/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type App struct {
	logger   logger.Logger
	migrator *migrate.Migrate
}

func New(cfg *config.Config) (*App, error) {
	logger, err := logger.New(cfg.MigratorConfig.LogLevel)
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create logger", err)
		return nil, err
	}

	db, err := sql.Open("postgres", cfg.Infrastructure.Postgres.ConnStr)
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create database connection", err)
		return nil, err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create database driver", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+cfg.MigratorConfig.MigrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		logger.Error("App", "App.Start", "Failed to create migrator", err)
		return nil, err
	}

	return &App{
		logger:   logger,
		migrator: m,
	}, nil
}
