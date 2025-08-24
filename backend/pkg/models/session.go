package models

import (
	"mp1/pkg/l10n"
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Token  string `json:"token" groups:"public" gorm:"not null"`
	UserID uint   `json:"userID" groups:"public" gorm:"not null"`
}

type SessionRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SessionResponse struct {
	*Session
}

func (session *Session) Create(db *gorm.DB, request SessionRequest, jwtKey []byte) error {
	var user User
	if err := db.Where("login = ?", request.Login).Take(&user).Error; err != nil {
		return err
	}
	match, err := user.VerifyPassword(request.Password)
	if err != nil {
		return err
	}
	if !match {
		return l10n.WrongPasswordError
	}
	token, err := util.GenerateJWT(request.Login, []byte(request.Password), jwtKey)
	if err != nil {
		return err
	}
	session.UserID = user.ID
	session.Token = token
	result := db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
