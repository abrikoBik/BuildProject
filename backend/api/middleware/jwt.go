package middleware

import (
	"mp1/util"
	"net/http"
	"os"
	"strings"
)

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
		next.ServeHTTP(w, r)
	})
}
