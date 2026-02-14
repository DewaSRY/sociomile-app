package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func TicketRouter(r chi.Router) {
	ticketHandler := handlers.NewTicketHandler()

	r.Route("/tickets", func(r chi.Router) {
		r.Use(middleware.JWTAuth)

		r.Post("/", ticketHandler.CreateTicket)
		r.Get("/number/{number}", ticketHandler.GetTicketByNumber)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", ticketHandler.GetTicket)
			r.Put("/", ticketHandler.UpdateTicket)
			r.Delete("/", ticketHandler.DeleteTicket)
		})
	})
}
