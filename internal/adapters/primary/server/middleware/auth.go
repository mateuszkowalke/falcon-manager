package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"falcon-manager/internal/adapters/primary/server/config"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey struct {
	key string
}

var ContextUserKey = &ContextKey{key: "user"}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("auth middleware entry")
		authHeader := r.Header.Get("Authorization")
		tokenString, formatOk := strings.CutPrefix(authHeader, "Bearer ")
		if !formatOk {
			slog.Error("authorization header in incorrect format", "reason", "expected Bearer token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		secret := config.JWTSecret()
		token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) { return []byte(secret), nil })
		if err != nil {
			slog.Error("couldn't parse the jwt token", "reason", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := token.Claims.GetSubject()
		if err != nil {
			slog.Error("couldn't get subject from jwt token", "reason", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, user)
		slog.Info("auth middleware exit")
		next(w, r.WithContext(ctx))
	}
}
