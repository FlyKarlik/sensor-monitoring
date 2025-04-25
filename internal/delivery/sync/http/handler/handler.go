package handler

import (
	"sensor-monitoring/internal/usecase"
	"sensor-monitoring/pkg/logger"
)

type Handler struct {
	logger  logger.Logger
	usecase *usecase.Usecase
}

func New(logger logger.Logger, usecase *usecase.Usecase) *Handler {
	return &Handler{
		logger:  logger,
		usecase: usecase,
	}
}
