package handler

import (
	"net/http"
	"sensor-monitoring/internal/delivery/sync/http/response"
	"sensor-monitoring/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchSensorData(c *gin.Context) {
	var input domain.SearchSensorDataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Error("handler", "SearchSensorData.ShouldBindJSON", "invalid request", err)
		response.ErrorResponse(c, "invalid request", http.StatusBadRequest)
		return
	}

	result := h.usecase.ISensorDataUsecase.SearchSensorData(c.Request.Context(), input)
	if !result.Success {
		h.logger.Error("handler", "SearchSensorData.SearchSensorData", "failed to search sensor data", result.Error)
		response.ErrorResponse(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)
}
