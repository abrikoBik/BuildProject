package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/liip/sheriff/v2"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func jsonDecode(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		http.Error(w, "bad json request", http.StatusBadRequest)
		return false
	}
	return true
}

func jsonEncode(w http.ResponseWriter, src interface{}, groups ...string) error {
	toEncode, err := sheriff.Marshal(&sheriff.Options{
		Groups: groups,
	}, src)
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
