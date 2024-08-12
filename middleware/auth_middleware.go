package middleware

import (
	"context"
	"net/http"
	"strings"
	"task-manager/utils"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Set the user information in the request context
		ctx := context.WithValue(r.Context(), "user", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
