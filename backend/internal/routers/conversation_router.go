package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func ConversationRouter(r chi.Router) {
	convHandler := handlers.NewConversationHandler()

	r.Route("/conversations", func(r chi.Router) {
		r.Use(middleware.JWTAuth)

		// Guest can create conversation
		r.Post("/", convHandler.CreateConversation)

		// Get user's own conversations
		r.Get("/my", convHandler.GetMyConversations)

		// Message routes
		r.Post("/messages", convHandler.CreateMessage)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", convHandler.GetConversation)
			r.Put("/status", convHandler.UpdateConversationStatus)
			r.Post("/assign", convHandler.AssignConversation)
			r.Get("/messages", convHandler.GetConversationMessages)
		})
	})
}
