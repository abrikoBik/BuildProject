package middleware

import (
	"context"
	// "mp1/models"
	"mp1/util"
	"net/http"
	"os"
	"strings"
	// "gorm.io/gorm"
)

const TokenKey = "JWT"

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header["Authorization"][0], " ")
		valid := strings.Compare(token[0], "Bearer") == 0
		valid = valid && util.ValidateJWT(token[1], os.Getenv("JWT_KEY"))
		// var tokens []string
		// db.Model(&models.UserDB{}).Pluck("sessions", &tokens)
		// valid = valid && db.Where("sessions IN ")
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), TokenKey, token[1])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
