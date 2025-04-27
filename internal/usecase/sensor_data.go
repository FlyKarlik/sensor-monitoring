package usecase

import (
	"context"
	"sensor-monitoring/internal/domain"
	"sensor-monitoring/internal/errs"
	"sensor-monitoring/internal/repository"
	"sensor-monitoring/pkg/generics"
	"sensor-monitoring/pkg/logger"
)

type sensorDataUsecase struct {
	logger logger.Logger
	repo   repository.ISensorDataRepository
}

func newSensorDataUsecase(repo repository.ISensorDataRepository, logger logger.Logger) ISensorDataUsecase {
	return &sensorDataUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *sensorDataUsecase) SearchSensorData(ctx context.Context, input domain.SearchSensorDataInput) generics.ItemsOutput[domain.SensorData] {
	total, err := u.repo.CountSensorData(ctx, input.SensorDataFilterInput)
	if err != nil {
		u.logger.Error("usecase", "SearchSensorData.CountSensorData", "failed to count sensor data", err)
		return generics.ItemsOutput[domain.SensorData]{
			Success: false,
			Error:   errs.ErrFailedToCountSensorData,
		}
	}

	if total == 0 {
		u.logger.Error("usecase", "SearchSensorData.CountSensorData", "no sensor data found", errs.ErrNoSensorData)
		return generics.ItemsOutput[domain.SensorData]{
			Success: false,
			Error:   errs.ErrNoSensorData,
		}
	}

	items, err := u.repo.SearchSensorData(ctx, input)
	if err != nil {
		u.logger.Error("usecase", "SearchSensorData.SearchSensorData", "failed to search sensor data", err)
		return generics.ItemsOutput[domain.SensorData]{
			Success: false,
			Error:   errs.ErrFailedToSearchSensorData,
		}
	}

	if len(items) == 0 {
		u.logger.Error("usecase", "SearchSensorData.SearchSensorData", "out of page limit", errs.ErrOutOfPageLimit)
		return generics.ItemsOutput[domain.SensorData]{
			Success: false,
			Error:   errs.ErrOutOfPageLimit,
		}
	}

	return generics.ItemsOutput[domain.SensorData]{
		Success: true,
		Total:   total,
		Items:   items,
	}
}
