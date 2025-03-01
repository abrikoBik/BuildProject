package api

import (
	"encoding/json"
	"fmt"
	"mp1/db"
	"mp1/util"
	"net/http"
)

// Creates a new database entry for user.
// Request body must be JSON.
func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user db.User
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("could not unmarshal json: "))
		w.Write([]byte(err.Error()))
		return
	}
	if user.Name == "" || user.Email == "" || user.PasswordHash == "" || user.Role == "" {
		w.WriteHeader(400)
		w.Write([]byte("invalid json request"))
		return
	}
	result := Database.Create(&user)
	if result.Error != nil {
		w.WriteHeader(500)
		w.Write([]byte("error creating user: "))
		w.Write([]byte(result.Error.Error()))
		return
	}
	util.Log("api", fmt.Sprintf("createUser: %d rows affected", result.RowsAffected))
	w.WriteHeader(200)
}
