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
		user models.User
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
	token := r.Context().Value(middleware.TokenKey).(string)
	// sanity checks
	var dbSession models.Session
	result := Database.Where("token = ?", token).Take(&dbSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user models.User
	result = Database.Where("name = ?", dbSession.Name).Take(&user)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/userInfo: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error finding user"))
		return
	}
	toMarshal := models.UserOut{
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
