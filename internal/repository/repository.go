package repository

import (
	"context"
	"sensor-monitoring/internal/domain"
	"sensor-monitoring/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ISensorDataRepository interface {
	CountSensorData(ctx context.Context, filter domain.SensorDataFilterInput) (int64, error)
	SearchSensorData(ctx context.Context, input domain.SearchSensorDataInput) ([]domain.SensorData, error)
}

type Repository struct {
	ISensorDataRepository
}

func New(db *pgxpool.Pool, logger logger.Logger) *Repository {
	return &Repository{
		ISensorDataRepository: newSensorDataRepository(db, logger),
	}
}
