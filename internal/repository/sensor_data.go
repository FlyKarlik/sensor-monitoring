package repository

import (
	"context"
	"sensor-monitoring/internal/domain"
	"sensor-monitoring/internal/repository/dao"
	"sensor-monitoring/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type sensorDataRepository struct {
	logger logger.Logger
	db     *pgxpool.Pool
}

func newSensorDataRepository(db *pgxpool.Pool, logger logger.Logger) ISensorDataRepository {
	return &sensorDataRepository{
		logger: logger,
		db:     db,
	}
}

func (r *sensorDataRepository) CountSensorData(ctx context.Context, filter domain.SensorDataFilterInput) (int64, error) {
	filterDAO := new(dao.SensorDataFilterDAO).FromSensorDataFilterInput(&filter)
	rows, err := r.db.Query(ctx,
		countSensorData,
		filterDAO.InferredBrand,
		filterDAO.Attestation,
		filterDAO.HasRecording,
		filterDAO.MinLengthTranscript,
	)
	if err != nil {
		r.logger.Error("repository", "CountSensorData.Query", "error counting sensor data", err)
		return 0, err
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			r.logger.Error("repository", "CountSensorData.Scan", "error scanning sensor data", err)
			return 0, err
		}
	}

	return count, nil
}

func (r *sensorDataRepository) SearchSensorData(ctx context.Context, input domain.SearchSensorDataInput) ([]domain.SensorData, error) {
	filterDAO := new(dao.SensorDataFilterDAO).FromSensorDataFilterInput(&input.SensorDataFilterInput)
	rows, err := r.db.Query(ctx,
		searchSensorData,
		filterDAO.InferredBrand,
		filterDAO.Attestation,
		filterDAO.HasRecording,
		filterDAO.MinLengthTranscript,
		filterDAO.SortInput.IsReverse,
		filterDAO.SortInput.Field,
		filterDAO.PaginationInput.Limit,
		filterDAO.PaginationInput.Page,
	)
	if err != nil {
		r.logger.Error("repository", "SearchSensorData.Query", "error searching sensor data", err)
		return nil, err
	}
	defer rows.Close()

	sensors := make([]domain.SensorData, 0)

	for rows.Next() {
		var sensor dao.SensorDataDAO
		err = rows.Scan(
			&sensor.ID,
			&sensor.CreatedAt,
			&sensor.UpdatedAt,
			&sensor.DeletedAt,
			&sensor.Timestamp,
			&sensor.OrginatingNumber,
			&sensor.SensorType,
			&sensor.Transcript,
			&sensor.RecordingFile,
			&sensor.StirShakenIdentityToken,
			&sensor.Attestation,
			&sensor.CertificateUrl,
			&sensor.SpC,
			&sensor.InferredViolation,
			&sensor.ShakenFailed,
			&sensor.UserId,
			&sensor.NoCert,
			&sensor.FlaggedBy,
			&sensor.Flagged,
			&sensor.CallbackTn,
			&sensor.RecordingInbound,
			&sensor.RecordingOutbound,
			&sensor.InferredBrand,
		)
		if err != nil {
			r.logger.Error("repository", "SearchSensorData.Scan", "error scanning sensor data", err)
			return nil, err
		}
		sensors = append(sensors, *sensor.FromDAO())
	}

	return sensors, nil
}
