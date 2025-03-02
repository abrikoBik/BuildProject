package api

import (
	"encoding/json"
	"fmt"
	"mp1/db"
	"mp1/util"
	"net/http"
)

// Creates a new database entry for user.
func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user db.User
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not unmarshal json: "))
		w.Write([]byte(err.Error()))
		return
	}

	result := Database.Create(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating user: "))
		w.Write([]byte(result.Error.Error()))
		return
	}
	util.Log("api", fmt.Sprintf("/createUser: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
}
