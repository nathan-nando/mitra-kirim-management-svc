package rest

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/config"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	logger *logrus.Logger
	app    *echo.Echo
	config *config.Config
}

func New() (*Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	logger := NewLogger()
	db := config.NewDatabase(cfg, logger)
	app := echo.New()

	config.BuildInternal(&config.AppConfig{
		Db:  db,
		App: app,
		Log: logger,
	})

	s := Server{
		logger: logger,
		app:    app,
		config: cfg,
	}

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	defer signal.Stop(stopServer)

	serverErrors := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		serverErrors <- s.app.Start(fmt.Sprintf(":%s", s.config.AppPort))
	}(&wg)

	// blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting REST API server: %w", err)
	case <-stopServer:
		s.logger.Warn("server received STOP signal")
		// asking listener to shutdown
		err := s.app.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("graceful shutdown did not complete: %w", err)
		}
		wg.Wait()
		s.logger.Info("server was shut down gracefully")
	}
	return nil
}
