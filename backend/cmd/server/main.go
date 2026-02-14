package main

import (
	_ "DewaSRY/sociomile-app/docs"

	"DewaSRY/sociomile-app/internal/config"
	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/routers"
	serviceImpl "DewaSRY/sociomile-app/internal/services/impl"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
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

func main() {
	// setup
	cfg := config.Load()
	logger.Init()

	db := database.Connect()

	// app context
	jwtSvc := jwtUtils.NewJwtService()
	authServiceSvc := serviceImpl.NewAuthService(jwtSvc)

	authorizeSvc := serviceImpl.NewAuthorizeService(db)
	hubSvc := serviceImpl.NewHubServiceImpl(db)

	conversationSvc := serviceImpl.NewConversationService()
	conversationMessageSvc := serviceImpl.NewConversationMessageService()
	organizationCrudSvc := serviceImpl.NewOrganizationCrudService()
	tickerSvc := serviceImpl.NewTicketService()
	organizationSvc := serviceImpl.NewOrganizationService(db)

	authHandler := handlers.NewAuthHandler(authServiceSvc)
	conversationHandler := handlers.NewConversationHandler(conversationSvc, conversationMessageSvc)
	organizationHandler := handlers.NewOrganizationHandler( organizationCrudSvc)
	orgStaffHandler := handlers.NewOrganizationStaffHandler( jwtSvc, organizationSvc)
	ticketHandler := handlers.NewTicketHandler(tickerSvc)

	hubHandler := handlers.NewHubHandler(hubSvc)

	authRouter := routers.AuthRouter{
		JwtService:  jwtSvc,
		AuthHandler: *authHandler,
	}

	conversationRouter := routers.ConversationRouter{
		JwtService:          jwtSvc,
		ConversationHandler: *conversationHandler,
	}

	organizationRouter := routers.OrganizationRouter{
		JwtService:          jwtSvc,
		AuthorizeService: authorizeSvc,
		OrganizationHandler: *organizationHandler,
		OrgStaffHandler: *orgStaffHandler,
	}

	tickerRouter := routers.TicketRouter{
		JwtService:    jwtSvc,
		TicketHandler: *ticketHandler,
	}

	hubRouter := routers.HubRouter{
		JwtService:       jwtSvc,
		AuthorizeService: authorizeSvc,
		HubHandler:       *hubHandler,
	}

	restAPIConfig := &config.RestAPIConfig{
		Config:             cfg,
		AuthRouter:         authRouter,
		HubRouter:          hubRouter,
		ConversationRouter: conversationRouter,
		OrganizationRouter: organizationRouter,
		TicketRouter:       tickerRouter,
	}

	restAPIConfig.Run()
}
