package handler

import (
	"net/http"
	"sensor-monitoring/internal/errs"
)

func getStatusFromError(err error) int {
	switch err {
	case errs.ErrNoSensorData:
		return http.StatusNotFound
	case errs.ErrFailedToCountSensorData:
		return http.StatusInternalServerError
	case errs.ErrFailedToSearchSensorData:
		return http.StatusInternalServerError
	case errs.ErrOutOfPageLimit:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
