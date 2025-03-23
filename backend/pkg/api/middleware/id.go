package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func ParseID(key, param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idParam := chi.URLParam(r, param)
			if strings.Compare(idParam, "") == 0 {
				http.Error(w, "missing id", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(idParam)
			if err != nil {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}
			withID := context.WithValue(r.Context(), key, id)
			next.ServeHTTP(w, r.WithContext(withID))
		})
	}
}
