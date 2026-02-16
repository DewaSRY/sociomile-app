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
// @BasePath        /api/v1
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
	authServiceSvc := serviceImpl.NewAuthService(db,jwtSvc)

	authorizeSvc := serviceImpl.NewAuthorizeService(db)
	hubSvc := serviceImpl.NewHubServiceImpl(db)

	organizationConversationSvc := serviceImpl.NewConversationService(db)
	organizationCrudSvc := serviceImpl.NewOrganizationCrudService(db)
	tickerSvc := serviceImpl.NewTicketService(db)
	organizationSvc := serviceImpl.NewOrganizationService(db)
	
	guestConversationSvc := serviceImpl.NewGuestConversationService(db)
	guestMessageSvc := serviceImpl.NewGuestMessageService(db)

	webHookSvc := serviceImpl.NewWebHookConversationService(db)
	
	authHandler := handlers.NewAuthHandler(authServiceSvc, jwtSvc)
	organizationHandler := handlers.NewOrganizationHandler(organizationCrudSvc)
	orgStaffHandler := handlers.NewOrganizationStaffHandler(jwtSvc, organizationSvc)
	organizationTicketHandler := handlers.NewOrganizationTicketHandler(tickerSvc)
	OrganizationConversationHandler := handlers.NewOrganizationConversationHandler(jwtSvc, organizationConversationSvc)

	hubHandler := handlers.NewHubHandler(hubSvc)
	guestConversationHandler := handlers.NewGuestConversationHandler(jwtSvc, guestConversationSvc)
	guestMessageHandler := handlers.NewGuestMessageHandler(jwtSvc, guestMessageSvc)

	webHookHandler := handlers.NewWebHookHandler(webHookSvc)

	authRouter := routers.AuthRouter{
		JwtService:  jwtSvc,
		AuthHandler: *authHandler,
	}

	organizationRouter := routers.OrganizationRouter{
		JwtService:          jwtSvc,
		AuthorizeService:    authorizeSvc,
		OrganizationHandler: *organizationHandler,
		OrgStaffHandler:     *orgStaffHandler,
		OrgTicketHandler: *organizationTicketHandler,
		OrgConversationHandler: *OrganizationConversationHandler,
	}

	hubRouter := routers.HubRouter{
		JwtService:       jwtSvc,
		AuthorizeService: authorizeSvc,
		HubHandler:       *hubHandler,
	}

	guestRoute := routers.GuestRouter{
		JwtService:               jwtSvc,
		GuestConversationHandler: *guestConversationHandler,
		GuestMessageHandler:      *guestMessageHandler,
	}

	webHookRoute := routers.WebHook{
		WebHookHandler : *webHookHandler,
	}

	restAPIConfig := &config.RestAPIConfig{
		Config:             cfg,
		AuthRouter:         authRouter,
		HubRouter:          hubRouter,
		OrganizationRouter: organizationRouter,
		GuestRouter:        guestRoute,
		WebHookRouter: webHookRoute,
	}

	restAPIConfig.Run()
}
