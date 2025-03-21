package middleware

import (
	"context"
	"mp1/pkg/util"
	"net/http"
	"os"
	"strings"
)

const JWT_CONTEXT_KEY = "JWT"

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.SplitAfter(authHeader, " ")
		if len(token) != 2 {
			http.Error(w, "invalid authorization header", http.StatusBadRequest)
			return
		}
		util.Info("api", "%s %s", token[0], token[1])
		if token[0] != "Bearer" || util.ValidateJWT(token[1], os.Getenv("JWT_KEY")) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		withToken := context.WithValue(r.Context(), JWT_CONTEXT_KEY, token)
		next.ServeHTTP(w, r.WithContext(withToken))
	})
}
