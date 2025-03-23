package models

import (
	"bytes"
	"mp1/pkg/util"
	"os"

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

func (session *Session) Create(db *gorm.DB, request SessionRequest) (err error) {
	var user User
	result := db.Where("login = ?", request.Login).Take(&user)
	if result.Error != nil {
		return result.Error
	}
	hash, err := util.GenerateHash([]byte(request.Password), user.Salt)
	if err != nil {
		return err
	}
	if bytes.Compare(user.Hash, hash) != 0 {
		return util.WrongPasswordError
	}
	token, err := util.GenerateJWT(request.Login, []byte(request.Password), os.Getenv("JWT_KEY"))
	if err != nil {
		return err
	}
	session.UserID = user.ID
	session.Token = token
	result = db.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
