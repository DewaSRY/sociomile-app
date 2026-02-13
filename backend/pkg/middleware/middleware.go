package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/lib/logger"
)

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// JWTAuth middleware validates JWT tokens and adds user info to context
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.ErrorLog("Missing authorization header", map[string]any{
				"path": r.URL.Path,
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Unauthorized",
				"message": "Missing authorization header",
			})
			return
		}

		// Extract the token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.ErrorLog("Invalid authorization header format", map[string]any{
				"path": r.URL.Path,
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format",
			})
			return
		}

		tokenString := parts[1]

		// Validate the token
		claims, err := jwtutil.ValidateToken(tokenString)
		if err != nil {
			logger.ErrorLog("Invalid or expired token", map[string]any{
				"path":  r.URL.Path,
				"error": err.Error(),
			})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
