package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"sensor-monitoring/internal/config"
	"sensor-monitoring/internal/delivery/sync/http/handler"
	"sensor-monitoring/pkg/logger"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(logger logger.Logger, cfg *config.Config, h *handler.Handler) *HTTPServer {
	router := initRouter(logger, cfg, h)
	httpServer := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.ServiceHost, cfg.ServicePort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		server: httpServer,
	}
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
