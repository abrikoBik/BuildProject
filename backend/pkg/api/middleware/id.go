package middleware

import (
	"context"
	"mp1/pkg/l10n"
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
				http.Error(w, l10n.MissingID, http.StatusBadRequest)
				return
			}
			id, err := strconv.ParseUint(idParam, 10, 32)
			if err != nil {
				http.Error(w, l10n.InvalidID, http.StatusBadRequest)
				return
			}
			withID := context.WithValue(r.Context(), key, uint(id))
			next.ServeHTTP(w, r.WithContext(withID))
		})
	}
}
