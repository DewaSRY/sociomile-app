package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)

type OrganizationRouter struct{
	JwtService	jwtUtils.JwtService
	OrganizationHandler handlers.OrganizationHandler
}

func (t *OrganizationRouter) Register(r chi.Router) {

	r.Route("/organizations", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))

		r.Post("/", t.OrganizationHandler.CreateOrganization)
		r.Get("/", t.OrganizationHandler.GetAllOrganizations)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", t.OrganizationHandler.GetOrganization)
			r.Put("/", t.OrganizationHandler.UpdateOrganization)
			r.Delete("/", t.OrganizationHandler.DeleteOrganization)
			r.Get("/stats", t.OrganizationHandler.GetOrganizationStats)
		})
	})
}
