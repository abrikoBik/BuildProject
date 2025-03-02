package middleware

import (
	"context"
	"mp1/util"
	"net/http"
	"os"
	"strings"
)

const TokenKey = "JWT"

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header["Authorization"][0], " ")
		if strings.Compare(token[0], "Bearer") != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !util.ValidateJWT(token[1], os.Getenv("JWT_KEY")) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), TokenKey, token[1])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
