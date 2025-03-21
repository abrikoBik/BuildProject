package handlers

import (
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
)

// Creates a new database entry for user.
func Register(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	user := models.User{
		Login: request.Login,
		Name:  request.Name,
		Email: request.Email,
	}
	if err := user.Create(db.Database, request); err != nil {
		util.Error("api:register", err)
		http.Error(w, "couldn't register user", http.StatusInternalServerError)
		return
	}
	util.Info("api:register", "%s registered", user.Name)
	w.WriteHeader(http.StatusCreated)
}
