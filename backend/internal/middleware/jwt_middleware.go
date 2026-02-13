package middleware

import (
	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
	"DewaSRY/sociomile-app/pkg/utils"
	"context"
	"net/http"
	"strings"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errorData := map[string]any{
				"path":    r.URL.Path,
				"error":   "Unauthorized",
				"message": "Missing authorization header",
			}
			logger.ErrorLog("Missing authorization header", errorData)
			utils.WriteJSONResponse(w, http.StatusUnauthorized, errorData)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			errorData := map[string]any{
				"path":    r.URL.Path,
				"error":   "Unauthorized",
				"message": "Missing authorization header",
			}
			logger.ErrorLog("Missing authorization header", errorData)
			utils.WriteJSONResponse(w, http.StatusUnauthorized, errorData)
			return
		}

		tokenString := parts[1]

		jwtInstance := jwtutil.InstanceJwtService()
		claims, err := jwtInstance.ValidateToken(tokenString)
		if err != nil {
			errorData := map[string]any{
				"path":    r.URL.Path,
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			}
			logger.ErrorLog("Invalid or expired token", errorData)
			utils.WriteJSONResponse(w, http.StatusUnauthorized, errorData)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)
		ctx = context.WithValue(ctx, "role_id", claims.RoleID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
