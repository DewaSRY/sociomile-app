package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(r chi.Router) {
	authHandler := handlers.NewAuthHandler()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.RefreshToken)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth)
			r.Get("/profile", authHandler.GetProfile)
		})
	})
}
