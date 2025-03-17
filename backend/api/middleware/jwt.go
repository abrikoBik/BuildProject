package middleware

import (
	"mp1/util"
	"net/http"
	"os"
	"strings"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.SplitAfter(authHeader, " ")
		if len(token) != 2 {
			http.Error(w, "invalid authorization header", http.StatusBadRequest)
			return
		}
		if token[0] != "Bearer" || util.ValidateJWT(token[1], os.Getenv("JWT_KEY")) {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
