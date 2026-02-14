package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)

type ConversationRouter struct {
	JwtService          jwtUtils.JwtService
	ConversationHandler handlers.ConversationHandler
}

func (t *ConversationRouter) Register(r chi.Router) {
	r.Route("/conversations", func(r chi.Router) {
		r.Use(middleware.JWTAuth(t.JwtService))

		// Guest can create conversation
		r.Post("/", t.ConversationHandler.CreateConversation)

		// Get user's own conversations
		r.Get("/my", t.ConversationHandler.GetMyConversations)

		// Message routes
		r.Post("/messages", t.ConversationHandler.CreateMessage)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", t.ConversationHandler.GetConversation)
			r.Put("/status", t.ConversationHandler.UpdateConversationStatus)
			r.Post("/assign", t.ConversationHandler.AssignConversation)
			r.Get("/messages", t.ConversationHandler.GetConversationMessages)
		})
	})
}
