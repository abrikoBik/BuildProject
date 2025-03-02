package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"mp1/db"
	"mp1/util"
	"net/http"

	"gorm.io/gorm"
)

// Creates a new database entry for user.
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var (
		user db.User
		err  error
	)
	if err := decoder.Decode(&user); err != nil {
		util.Log("api", fmt.Sprintf("/register: %s", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Salt, err = util.RandomBytes(128)
	if err != nil {
		util.Log("api", fmt.Sprintf("/register: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	user.Hash, err = util.GenerateHash([]byte(user.Password), user.Salt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/register: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	result := Database.Create(&user)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/register: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating user"))
		return
	}
	util.Log("api", fmt.Sprintf("/register: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var session db.Session
	if err := decoder.Decode(&session); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// sanity checks
	var dbSession db.Session
	result := Database.Where("token = ?", session.Token).Take(&dbSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user db.User
	result = Database.Where("name = ?", dbSession.Name).Take(&user)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/userInfo: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error finding user"))
		return
	}
	toMarshal := db.UserOut{
		Name:     user.Name,
		FullName: user.FullName,
		Role:     user.Role,
		ID:       user.ID,
	}
	marshalled, err := json.Marshal(toMarshal)
	if err != nil {
		util.Log("api", fmt.Sprintf("/userInfo: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshalling user"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalled)
}
