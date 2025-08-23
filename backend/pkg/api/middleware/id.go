package middleware

import (
	"context"
	"mp1/pkg/l10n"
	"mp1/pkg/util"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func ParseID(key ContextKey, param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idParam := chi.URLParam(r, param)
			if len(idParam) == 0 {
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

func GetID(key ContextKey, context context.Context) (uint, error) {
	id, ok := context.Value(key).(uint)
	if !ok {
		return 0, util.InvalidIDError
	}
	return id, nil
}

func ParseUUID(key ContextKey, param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idParam := chi.URLParam(r, param)
			if len(idParam) == 0 {
				http.Error(w, l10n.MissingID, http.StatusBadRequest)
				return
			}
			id, err := uuid.Parse(idParam)
			if err != nil {
				http.Error(w, l10n.InvalidID, http.StatusBadRequest)
				return
			}
			withID := context.WithValue(r.Context(), key, id)
			next.ServeHTTP(w, r.WithContext(withID))
		})
	}
}

func GetUUID(key ContextKey, context context.Context) (uuid.UUID, error) {
	id, ok := context.Value(key).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, util.InvalidIDError
	}
	return id, nil
}
