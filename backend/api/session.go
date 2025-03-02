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
		w.Write([]byte("could not unmarshal json: "))
		w.Write([]byte(err.Error()))
		return
	}

	// sanity checks
	var dbUser db.User
	result := Database.Where("name = ?", user.Name).Take(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not find user"))
		return
	}
	hash := util.GenerateHash([]byte(user.Password), dbUser.Salt)
	if bytes.Compare(dbUser.Hash, hash) != 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":  "mp1-auth",
		"sub":  user.Name,
		"hash": dbUser.Hash,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		util.Log("api", fmt.Sprintf("/createSession: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session := db.Session{
		Name:      user.Name,
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7 * 5), // 5 weeks
	}
	result = Database.Create(&session)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating session: "))
		w.Write([]byte(result.Error.Error()))
		return
	}
	util.Log("api", fmt.Sprintf("/createSession: %d row(s) affected", result.RowsAffected))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(tokenStr))
}
