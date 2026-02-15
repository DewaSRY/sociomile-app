package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)

type GuestRouter struct {
	JwtService               jwtUtils.JwtService
	GuestConversationHandler handlers.GuestConversationHandler
}

func (t *GuestRouter) Register(r chi.Router) {
	r.Route("/guest", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))

		r.Route("/conversations", func(r chi.Router) {
			r.Get("/", t.GuestConversationHandler.GetConversation)
			r.Post("/", t.GuestConversationHandler.CreateConversation)
		})

	})
}
