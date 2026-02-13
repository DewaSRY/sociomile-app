package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/routers"
	"DewaSRY/sociomile-app/pkg/lib/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.InfoLog("shutting down gracefully, press Ctrl+C again to force", map[string]any{})
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

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	logger.Init()
	database.Connect()

	// // Auto-migrate database models
	// if err := database.DB.AutoMigrate(&models.User{}); err != nil {
	// 	logger.ErrorLog("Failed to auto-migrate database", map[string]any{
	// 		"error": err.Error(),
	// 	})
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	// logger.InfoLog("Database migration completed", map[string]any{})

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.AllowContentType("application/json"))

	r.Route("/api/v1", func(r chi.Router) {
		// Register authentication routes
		routers.AuthRouter(r)
	})


	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	done := make(chan bool, 1)
	go gracefulShutdown(server, done)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.ErrorLog("http server error", map[string]any{
			"server": err.Error(),
		})
	}

	<-done
	log.Println("Graceful shutdown complete.")
	logger.LoggerSync()
}
