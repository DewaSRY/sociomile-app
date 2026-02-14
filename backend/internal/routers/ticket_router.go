package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)

type TicketRouter struct{
	JwtService jwtUtils.JwtService
	TicketHandler handlers.TicketHandler
}

func (t*TicketRouter) Register (r chi.Router) {

	r.Route("/tickets", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))

		r.Post("/", t.TicketHandler.CreateTicket)
		r.Get("/number/{number}", t.TicketHandler.GetTicketByNumber)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", t.TicketHandler.GetTicket)
			r.Put("/", t.TicketHandler.UpdateTicket)
			r.Delete("/", t.TicketHandler.DeleteTicket)
		})
	})
}
