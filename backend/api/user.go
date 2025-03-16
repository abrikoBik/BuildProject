package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"mp1/api/middleware"
	"mp1/models"
	"mp1/util"
	"net/http"

	"gorm.io/gorm"
)

// Creates a new database entry for user.
func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var (
		userIn models.UserJSONIn
		err    error
	)
	if err := decoder.Decode(&userIn); err != nil {
		http.Error(w, "bad json request", http.StatusBadRequest)
		return
	}

	userDB := models.User{
		Name:        userIn.Name,
		FullName:    userIn.FullName,
		Email:       userIn.Email,
		PhoneNumber: userIn.PhoneNumber,
	}
	userDB.Salt, err = util.RandomBytes(128)
	if err != nil {
		util.Log("api", fmt.Sprintf("/register: %s", err))
		http.Error(w, "couldn't generate salt", http.StatusInternalServerError)
		return
	}
	userDB.Hash, err = util.GenerateHash([]byte(userIn.Password), userDB.Salt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/register: %s", err))
		http.Error(w, "couldn't generate hash", http.StatusInternalServerError)
		return
	}
	result := Database.Create(&userDB)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/register: %s", result.Error.Error()))
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	util.Log("api", fmt.Sprintf("/register: new user %s", userDB.Name))
	w.WriteHeader(http.StatusCreated)
}

// TODO: rewrite
func userInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.TokenKey).(string)
	// sanity checks
	var dbSession models.Session
	result := Database.Where("token = ?", token).Take(&dbSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user models.User
	result = Database.Where("sessions IN ?", dbSession.Token).Take(&user)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/userInfo: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error finding user"))
		return
	}
	var toMarshal models.UserJSONOut
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
