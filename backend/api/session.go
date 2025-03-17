package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
	"os"

	"github.com/liip/sheriff/v2"
	"gorm.io/gorm"
)

// Creates a new database entry for session.
func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request models.SessionRequest
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "bad json request", http.StatusBadRequest)
		return
	}

	// sanity checks
	var dbUser models.User
	result := db.Database.Where("name = ?", request.Name).Take(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	hash, err := util.GenerateHash([]byte(request.Password), dbUser.Salt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err))
		http.Error(w, "couldn't generate hash", http.StatusInternalServerError)
		return
	}
	if bytes.Compare(dbUser.Hash, hash) != 0 {
		http.Error(w, "wrong password", http.StatusForbidden)
		return
	}
	token, err := util.GenerateJWT(request.Name, []byte(request.Password), os.Getenv("JWT_KEY"))
	if err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err.Error()))
		http.Error(w, "couldn't generate token", http.StatusInternalServerError)
		return
	}
	session := models.Session{
		UserID: dbUser.ID,
		Token:  token,
	}
	result = db.Database.Create(&session)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/login: %s", result.Error.Error()))
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	response := models.SessionResponse{
		Session: &models.Session{
			UserID: dbUser.ID,
			Token:  token,
		},
	}
	toEncode, err := sheriff.Marshal(&sheriff.Options{
		Groups: []string{"public"},
	}, response)
	if err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err.Error()))
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	if err = encoder.Encode(toEncode); err != nil {
		util.Log("api", fmt.Sprintf("/login: %s", err.Error()))
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return
	}
	util.Log("api", fmt.Sprintf("/login: %s logged in", dbUser.Name))
	w.WriteHeader(http.StatusCreated)
}

// Removes the database entry for a session.
// TODO: rewrite
func logout(w http.ResponseWriter, r *http.Request) {
	// sanity checks
	// var dbSession models.Session
	// result := db.Database.Where("token = ?", token).Take(&dbSession)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	// result = db.Database.Where("token = ?", token).Delete(&dbSession)
	// if result.Error != nil {
	// 	util.Log("api", fmt.Sprintf("/logout: %s", result.Error.Error()))
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("error invalidating session"))
	// 	return
	// }
	// util.Log("api", fmt.Sprintf("/logout: %d row(s) affected", result.RowsAffected))
	// w.WriteHeader(http.StatusOK)
}
