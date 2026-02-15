package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type WebHook struct {
	WebHookHandler       handlers.WebHookHandler
}

func (t *WebHook) Register(r chi.Router) {
	r.Route("/webhooks", func(r chi.Router) {
		r.Post("/conversations", t.WebHookHandler.CreateConversation)
	})
}
