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
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Creates a new database entry for session.
func createSession(w http.ResponseWriter, r *http.Request) {
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
		util.Log("api", fmt.Sprintf("/createSession: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	if bytes.Compare(dbUser.Hash, hash) != 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userSalt, err := util.RandomBytes(64)
	if err != nil {
		util.Log("api", fmt.Sprintf("/createSession: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	userHash, err := util.GenerateHash([]byte(user.Password), userSalt)
	if err != nil {
		util.Log("api", fmt.Sprintf("/createSession: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("hash error"))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":  "mp1-auth",
		"sub":  user.Name,
		"hash": userHash,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		util.Log("api", fmt.Sprintf("/createSession: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating token"))
		return
	}

	session := db.Session{
		Name:      user.Name,
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 5), // 5 weeks
	}
	result = Database.Create(&session)
	if result.Error != nil {
		util.Log("api", fmt.Sprintf("/createSession: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating session"))
		return
	}
	util.Log("api", fmt.Sprintf("/createSession: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(tokenStr))
}

// Removes the database entry for a session.
func invalidateSession(w http.ResponseWriter, r *http.Request) {
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
		util.Log("api", fmt.Sprintf("/invalidateSession: %s", result.Error.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error invalidating session"))
		return
	}
	util.Log("api", fmt.Sprintf("/invalidateSession: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusOK)
}
