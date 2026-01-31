package router

import (
	"context"
	"net/http"
	"strings"
	"url-shortener/auth"
)

func AuthMiddleware(ctx context.Context, authService auth.Auth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error": "токен отсутствует"}`, http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, `{"error": "неверный формат токена"}`, http.StatusUnauthorized)
				return
			}
			tokenString := parts[1]

			userID, _, err := authService.ValidateAccessToken(ctx, tokenString)
			if err != nil {
				http.Error(w, `{"error": "невалидный токен"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
