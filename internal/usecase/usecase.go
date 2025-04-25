package usecase

import (
	"context"
	"sensor-monitoring/internal/domain"
	"sensor-monitoring/internal/repository"
	"sensor-monitoring/pkg/generics"
	"sensor-monitoring/pkg/logger"
)

type ISensorDataUsecase interface {
	SearchSensorData(ctx context.Context, input domain.SearchSensorDataInput) generics.ItemsOutput[domain.SensorData]
}

type Usecase struct {
	ISensorDataUsecase
}

func New(repo *repository.Repository, logger logger.Logger) *Usecase {
	return &Usecase{
		ISensorDataUsecase: newSensorDataUsecase(repo, logger),
	}
}
