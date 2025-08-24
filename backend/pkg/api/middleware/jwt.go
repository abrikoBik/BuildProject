package middleware

import (
	"context"
	"mp1/pkg/l10n"
	"mp1/pkg/util"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(jwtKey []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			token := strings.SplitAfter(authHeader, " ")
			if len(token) != 2 || strings.TrimSpace(token[0]) != "Bearer" {
				http.Error(w, l10n.InvalidAuthHeader, http.StatusBadRequest)
				return
			}
			validated := util.ValidateJWT(strings.TrimSpace(token[1]), jwtKey)
			if !validated.Valid {
				http.Error(w, l10n.InvalidToken, http.StatusUnauthorized)
				return
			}
			withToken := context.WithValue(r.Context(), JWTKey, validated)
			next.ServeHTTP(w, r.WithContext(withToken))
		})
	}
}

func GetToken(context context.Context) (jwt.Token, error) {
	token, ok := context.Value(JWTKey).(jwt.Token)
	if !ok {
		return jwt.Token{}, l10n.InvalidTokenError
	}
	return token, nil
}
