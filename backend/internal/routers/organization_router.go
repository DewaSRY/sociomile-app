package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)

func OrganizationRouter(r chi.Router, jwtService jwtUtils.JwtService) {
	orgHandler := handlers.NewOrganizationHandler()

	r.Route("/organizations", func(r chi.Router) {
		r.Use(middleware.JWTAuth(jwtService))

		r.Post("/", orgHandler.CreateOrganization)
		r.Get("/", orgHandler.GetAllOrganizations)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", orgHandler.GetOrganization)
			r.Put("/", orgHandler.UpdateOrganization)
			r.Delete("/", orgHandler.DeleteOrganization)
			r.Get("/stats", orgHandler.GetOrganizationStats)
		})
	})
}
