package middleware

import (
	"net/http"
	"sensor-monitoring/internal/config"
	"sensor-monitoring/internal/delivery/sync/http/response"
	"sensor-monitoring/internal/errs"
	"sensor-monitoring/pkg/logger"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(logger logger.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authKey := c.GetHeader("X-Auth-Key")
		if authKey == "" {
			logger.Error("middleware", "AuthMiddleware", errs.ErrEmptyAuthKey.Error(), errs.ErrEmptyAuthKey)
			response.ErrorResponse(c, errs.ErrEmptyAuthKey.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		if authKey != cfg.AuthKey {
			logger.Error("middleware", "AuthMiddleware", errs.ErrInvalidAuthKey.Error(), errs.ErrInvalidAuthKey)
			response.ErrorResponse(c, errs.ErrInvalidAuthKey.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
