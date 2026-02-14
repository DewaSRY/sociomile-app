package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "DewaSRY/sociomile-app/docs"

	"DewaSRY/sociomile-app/internal/config"
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/routers"
	serviceImpl "DewaSRY/sociomile-app/internal/services/impl"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Sociomile API
// @version         1.0
// @description     This is a Sociomile application server with authentication.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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

func main() {
	// setup 
	config := config.Load()
	logger.Init()
	db:=database.Connect()

	// app context 
	jwtServiceInstance := jwtUtils.NewJwtService()
	authServiceSvc := serviceImpl.NewAuthService(jwtServiceInstance)

	authorizeSvc := serviceImpl.NewAuthorizeService(db);
	hubSvc := serviceImpl.NewHubServiceImpl(db);

	conversationSvc :=  serviceImpl.NewConversationService();
	conversationMessageSvc :=  serviceImpl.NewConversationMessageService();
	organizationSvc := serviceImpl.NewOrganizationService();
	tickerSvc := serviceImpl.NewTicketService()


	authHandler := handlers.NewAuthHandler(authServiceSvc);
	conversationHandler := handlers.NewConversationHandler(conversationSvc, conversationMessageSvc);
	organizationHandler := handlers.NewOrganizationHandler(organizationSvc);
	ticketHandler:= handlers.NewTicketHandler(tickerSvc);

	hubHandler := handlers.NewHubHandler(hubSvc)
	
	authRouter:= routers.AuthRouter{
		JwtService: jwtServiceInstance,
		AuthHandler: *authHandler,
	}

	conversationRouter := routers.ConversationRouter{
		JwtService: jwtServiceInstance,
		ConversationHandler: *conversationHandler,
	}

	organizationRouter:= routers.OrganizationRouter{
		JwtService: jwtServiceInstance,
		OrganizationHandler: *organizationHandler,
	}

	tickerRouter:= routers.TicketRouter{
		JwtService: jwtServiceInstance,
		TicketHandler: *ticketHandler,
	}

	hubRouter := routers.HubRouter{
		JwtService: jwtServiceInstance,
		AuthorizeService: authorizeSvc,
		HubHandler: *hubHandler,
	}

	// server register
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.AllowContentType("application/json"))

	// Swagger documentation
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", config.Host)),
	))

	r.Route("/api/v1", func(r chi.Router) {
		authRouter.Register(r)
		hubRouter.Register(r)

		conversationRouter.Register(r)
		organizationRouter.Register(r)
		tickerRouter.Register(r)
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	done := make(chan bool, 1)
	go gracefulShutdown(server, done)

	logger.InfoLog(fmt.Sprintf("json documentation:  %s/swagger/doc.json", config.Host), map[string]any{
		"message": fmt.Sprintf("Server running in : %s", config.Host),
	})

	logger.InfoLog(fmt.Sprintf("client documentation: %s/swagger/index.html", config.Host), map[string]any{
		"message": fmt.Sprintf("Server running in : %s", config.Host),
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
