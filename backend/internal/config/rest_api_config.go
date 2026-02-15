package config

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"DewaSRY/sociomile-app/internal/routers"
	"DewaSRY/sociomile-app/pkg/lib/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RestAPIConfig struct {
	Config             *Config
	AuthRouter         routers.AuthRouter
	HubRouter          routers.HubRouter
	OrganizationRouter routers.OrganizationRouter
	GuestRouter	routers.GuestRouter
	WebHookRouter routers.WebHook
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		logger.ErrorLog("Server forced to shutdown with error", map[string]any{
			"server": err.Error(),
		})
	}
	logger.InfoLog("Server exiting", map[string]any{})
	done <- true
}

func (cfg *RestAPIConfig) Run() {
	// server register
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.AllowContentType("application/json"))

	// Swagger documentation
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", cfg.Config.Host)),
	))

	r.Route("/api/v1", func(r chi.Router) {
		cfg.AuthRouter.Register(r)
		cfg.HubRouter.Register(r)
		cfg.OrganizationRouter.Register(r)
		cfg.GuestRouter.Register(r)
		cfg.WebHookRouter.Register(r)
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Config.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	done := make(chan bool, 1)
	go gracefulShutdown(server, done)

	logger.InfoLog(fmt.Sprintf("json documentation:  %s/swagger/doc.json", cfg.Config.Host), map[string]any{
		"message": fmt.Sprintf("Server running in : %s", cfg.Config.Host),
	})

	logger.InfoLog(fmt.Sprintf("client documentation: %s/swagger/index.html", cfg.Config.Host), map[string]any{
		"message": fmt.Sprintf("Server running in : %s", cfg.Config.Host),
	})

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.ErrorLog("http server error", map[string]any{
			"server": err.Error(),
		})
	}

	<-done
	logger.InfoLog("Graceful shutdown complete.", map[string]any{
		"message": "success",
	})
	logger.LoggerSync()
}