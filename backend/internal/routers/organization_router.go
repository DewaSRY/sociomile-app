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
	JwtService             jwtUtils.JwtService
	AuthorizeService       services.AuthorizeService
	OrganizationHandler    handlers.OrganizationHandler
	OrgStaffHandler        handlers.OrganizationStaffHandler
	OrgTicketHandler       handlers.OrganizationTicketHandler
	OrgConversationHandler handlers.OrganizationConversationHandler
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

		r.Route("/ticket", func(r chi.Router) {
			r.Get("/", t.OrgTicketHandler.GetTicketsList)
			r.Post("/", t.OrgTicketHandler.CreateTicket)
			r.Route("/{id}", func(r chi.Router) {
				r.Put("/", t.OrgTicketHandler.UpdateTicket)
			})
		})

		r.Route("/conversations", func(r chi.Router) {
			r.Get("/", t.OrgConversationHandler.GetConversationsList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", t.OrgConversationHandler.GetConversationByID)
				r.Put("/assign", t.OrgConversationHandler.AssignConversation)
				r.Put("/status", t.OrgConversationHandler.UpdateConversationStatus)
			})
		})
	})

}
