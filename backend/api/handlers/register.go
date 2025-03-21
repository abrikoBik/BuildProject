package handlers

import (
	"errors"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
	"net/mail"
	"strings"

	"gorm.io/gorm"
)

// Creates a new database entry for user.
func Register(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	if strings.Compare(request.Password, "") == 0 {
		http.Error(w, "empty password", http.StatusBadRequest)
		return
	}

	user := models.User{
		Login: request.Login,
		Name:  request.Name,
		Email: request.Email,
	}
	if err := user.Create(db.Database, request); err != nil {
		util.Error("api:register", err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			http.Error(w, "email or name taken", http.StatusBadRequest)
		} else {
			http.Error(w, "couldn't register user", http.StatusInternalServerError)
		}
		return
	}
	util.Info("api:register", "%s registered", user.Name)
	w.WriteHeader(http.StatusCreated)
}
