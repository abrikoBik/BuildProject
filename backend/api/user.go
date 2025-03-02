package api

import (
	"encoding/json"
	"fmt"
	"mp1/db"
	"mp1/util"
	"net/http"
)

// Creates a new database entry for user.
// Endpoint: /register
func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var (
		user db.User
		err  error
	)
	if err := decoder.Decode(&user); err != nil {
		util.Log("api", fmt.Sprintf("/createUser: %s", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Salt, err = util.RandomBytes(128)
	if err != nil {
		util.Log("api", fmt.Sprintf("/createUser: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	user.Hash, err = util.GenerateHash([]byte(user.Password), user.Salt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/createUser: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	result := Database.Create(&user)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/createUser: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating user"))
		return
	}
	util.Log("api", fmt.Sprintf("/createUser: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
}
