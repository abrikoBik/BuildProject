package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mp1/db"
	"mp1/util"
	"net/http"
	"os"

	"gorm.io/gorm"
)

// Creates a new database entry for session.
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user db.SessionUser
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// sanity checks
	var dbUser db.User
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

	session := db.Session{
		Name:  user.Name,
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
	if r.Method != "DELETE" {
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result = Database.Where("token = ?", session.Token).Delete(&dbSession)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/logout: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error invalidating session"))
		return
	}
	util.Log("api", fmt.Sprintf("/logout: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusOK)
}
