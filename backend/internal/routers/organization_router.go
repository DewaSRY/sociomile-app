package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"
	"DewaSRY/sociomile-app/internal/services"

	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"

	"github.com/go-chi/chi/v5"
)

type OrganizationRouter struct {
	JwtService          jwtUtils.JwtService
	AuthorizeService    services.AuthorizeService
	OrganizationHandler handlers.OrganizationHandler
	OrgStaffHandler handlers.OrganizationStaffHandler
}

func (t *OrganizationRouter) Register(r chi.Router) {

	r.Route("/organizations", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))

		r.Route("/staff", func(r chi.Router) {
			r.With(middleware.Authorize(
				t.JwtService,
				t.AuthorizeService,
				[]string{
					models.RoleOrganizationOwner,
				},
			)).Post("/", t.OrgStaffHandler.CreateOrganizationStaff)

			r.With(middleware.Authorize(
				t.JwtService,
				t.AuthorizeService,
				[]string{
					models.RoleOrganizationOwner,
					models.RoleOrganizationSales,
				},
			)).Get("/", t.OrgStaffHandler.GetStaffListPagination)
		})

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
