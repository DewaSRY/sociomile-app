package middleware

import (
	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func JWTAuth(jwtInstance jwtutil.JwtService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteJSONResponse(w, http.StatusUnauthorized, map[string]any{
					"path":    r.URL.Path,
					"error":   "unauthorized",
					"message": "missing authorization header",
				})
				return
			}

			parts := strings.Fields(authHeader)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				utils.WriteJSONResponse(w, http.StatusUnauthorized, map[string]any{
					"path":    r.URL.Path,
					"error":   "unauthorized",
					"message": "invalid authorization header format",
				})
				return
			}

			tokenString := parts[1]

			claims, err := jwtInstance.ValidateToken(tokenString)
			if err != nil {
				utils.WriteJSONResponse(w, http.StatusUnauthorized, map[string]any{
					"path":    r.URL.Path,
					"error":   "unauthorized",
					"message": "invalid or expired token",
				})
				return
			}

			ctx := context.WithValue(r.Context(), jwtutil.UserContextKey, claims)
			fmt.Println(" the calim store ", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
