package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"
	"DewaSRY/sociomile-app/internal/services"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"

	"github.com/go-chi/chi/v5"
)

type HubRouter struct {
	JwtService       jwtUtils.JwtService
	AuthorizeService services.AuthorizeService
	HubHandler       handlers.HubHandler
}

func (t *HubRouter) Register(r chi.Router) {
	r.Route("/hub", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))
		r.Use(middleware.Authorize(t.JwtService, t.AuthorizeService, []string{
			models.RoleSuperAdmin,
		}))

		r.Get("/organizations", t.HubHandler.GetOrganizationPagination)
		r.Post("/organizations", t.HubHandler.CreateOrganization)
	})
}
