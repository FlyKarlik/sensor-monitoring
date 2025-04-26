package httpserver

import (
	"net/http/pprof"
	"sensor-monitoring/internal/config"
	"sensor-monitoring/internal/delivery/sync/http/handler"
	"sensor-monitoring/internal/delivery/sync/http/middleware"
	"sensor-monitoring/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter(logger logger.Logger, cfg *config.Config, h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	router.GET("/health", h.HealthCheck)
	registerPprof(router)

	v1 := router.Group("/v1")
	{
		registerSensorDataRoutes(logger, cfg, v1, h)
	}

	return router
}

func registerSensorDataRoutes(logger logger.Logger, cfg *config.Config, router *gin.RouterGroup, h *handler.Handler) {
	sensorData := router.Group("/sensor-data")
	sensorData.Use(middleware.JSONMiddleware(), middleware.AuthMiddleware(logger, cfg))
	{
		sensorData.POST("/search", h.SearchSensorData)
	}
}

func registerPprof(router *gin.Engine) {
	pprofGroup := router.Group("/debug/pprof")
	{
		router.GET("/debug/vars", gin.WrapH(pprof.Handler("vars")))
		pprofGroup.GET("/", gin.WrapF(pprof.Index))
		pprofGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		pprofGroup.GET("/profile", gin.WrapF(pprof.Profile))
		pprofGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
		pprofGroup.GET("/trace", gin.WrapF(pprof.Trace))
		pprofGroup.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		pprofGroup.GET("/block", gin.WrapH(pprof.Handler("block")))
		pprofGroup.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		pprofGroup.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		pprofGroup.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		pprofGroup.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
