package handlers

import (
	"encoding/json"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"net/http"
	"os"

	"github.com/liip/sheriff/v2"
)

var jwtKey []byte

func init() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func jsonDecode(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		http.Error(w, "bad json request", http.StatusBadRequest)
		return false
	}
	return true
}

func jsonEncode(w http.ResponseWriter, src interface{}, groups ...string) error {
	toEncode, err := sheriff.Marshal(&sheriff.Options{Groups: groups}, src)
	if err != nil {
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return err
	}
	if err = json.NewEncoder(w).Encode(toEncode); err != nil {
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func verifyUser(user models.User, w http.ResponseWriter, r *http.Request) bool {
	token, err := middleware.GetToken(r.Context())
	if err != nil {
		http.Error(w, l10n.InvalidToken, http.StatusInternalServerError)
		return false
	}
	verified, err := user.VerifyToken(db.Database, token)
	if err != nil {
		http.Error(w, l10n.InvalidToken, http.StatusInternalServerError)
		return false
	}
	if !verified {
		http.Error(w, l10n.InvalidToken, http.StatusForbidden)
		return false
	}
	return true
}
