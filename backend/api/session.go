package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mp1/api/middleware"
	"mp1/models"
	"mp1/util"
	"net/http"
	"os"

	"gorm.io/gorm"
)

// Creates a new database entry for session.
func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.SessionJSONIn
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// sanity checks
	var dbUser models.UserDB
	result := Database.Where("name = ?", user.Name).Take(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash, err := util.GenerateHash([]byte(user.Password), dbUser.Salt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	if bytes.Compare(dbUser.Hash, hash) != 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	token, err := util.GenerateJWT(user.Name, []byte(user.Password), os.Getenv("JWT_KEY"))
	if err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating token"))
		return
	}

	session := models.Session{
		Token: token,
	}
	result = Database.Create(&session)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/login: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating session"))
		return
	}
	util.Log("api", fmt.Sprintf("/login: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(token))
}

// Removes the database entry for a session.
func logout(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.TokenKey).(string)
	// sanity checks
	var dbSession models.Session
	result := Database.Where("token = ?", token).Take(&dbSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result = Database.Where("token = ?", token).Delete(&dbSession)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/logout: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error invalidating session"))
		return
	}
	util.Log("api", fmt.Sprintf("/logout: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusOK)
}
