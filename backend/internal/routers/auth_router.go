package routers

import (
	"DewaSRY/sociomile-app/internal/handlers"
	"DewaSRY/sociomile-app/internal/middleware"
	jwtUtils "DewaSRY/sociomile-app/pkg/lib/jwt"

	"github.com/go-chi/chi/v5"
)


type AuthRouter struct{
	JwtService  jwtUtils.JwtService
	AuthHandler handlers.AuthHandler
}

func(t *AuthRouter) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", t.AuthHandler.Register)
		r.Post("/login", t.AuthHandler.Login)
		r.Post("/refresh", t.AuthHandler.RefreshToken)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth(t.JwtService))
			r.Get("/profile", t.AuthHandler.GetProfile)
		})
	})
}
