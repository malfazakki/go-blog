package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/malfazakki/go-blog/handlers"
	"github.com/malfazakki/go-blog/utils"
)

type ContextKey string

const userIDKey ContextKey = "userID"

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handlers.ErrorResponse(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		// Check if the Authorization header has the correct format
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			handlers.ErrorResponse(w, http.StatusUnauthorized, "Authorization header format is invalid (Bearer {token})")
			return
		}

		// Get the token
		tokenString := headerParts[1]

		// Validate the token
		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			handlers.ErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
