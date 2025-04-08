package models

import (
	"bytes"
	"mp1/pkg/util"
	"net/mail"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string    `json:"login" groups:"public" gorm:"unique;not null"`
	Name     string    `json:"name" groups:"public" gorm:"not null"`
	Email    string    `json:"email" groups:"public" gorm:"unique;not null"`
	Hash     []byte    `gorm:"not null"`
	Salt     []byte    `gorm:"not null"`
	Sessions []Session `gorm:"foreignKey:UserID"`
	Worker   Worker    `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

type UserResponse struct {
	*User
	RegisteredAt string `json:"registeredAt" groups:"public"`
	ID           uint   `json:"id" groups:"register"`
}

func (user *User) Create(db *gorm.DB, request UserRequest) (err error) {
	if _, err = mail.ParseAddress(request.Email); err != nil {
		return util.InvalidEmailError
	}
	user.Email = request.Email
	if len(request.Password) < util.MinPasswordLength {
		return util.InsecurePasswordError
	}
	if err = user.PreparePassword(request.Password); err != nil {
		return
	}
	user.Name = cases.Title(language.Russian).String(request.Name)
	user.Login = request.Login
	return db.Create(&user).Error
}

func (user *User) Patch(db *gorm.DB, request UserRequest) (err error) {
	if len(request.User.Email) > 0 {
		if _, err = mail.ParseAddress(request.User.Email); err != nil {
			return util.InvalidEmailError
		}
		user.Email = request.User.Email
	}
	if len(request.Password) > 0 {
		if len(request.Password) < util.MinPasswordLength {
			return util.InsecurePasswordError
		}
		if err = user.PreparePassword(request.Password); err != nil {
			return
		}
	}
	if len(request.User.Name) > 0 {
		user.Name = cases.Title(language.Russian).String(request.User.Name)
	}
	return db.Save(&user).Error
}

func (user *User) PreparePassword(password string) (err error) {
	if user.Salt, err = util.RandomBytes(16); err != nil {
		return
	}
	user.Hash, err = util.GenerateHash([]byte(password), user.Salt)
	return
}

func (user User) VerifyPassword(db *gorm.DB, password string) (bool, error) {
	hash, err := util.GenerateHash([]byte(password), user.Salt)
	if err != nil {
		return false, err
	}
	return (bytes.Compare(user.Hash, hash) == 0), nil
}

func (user User) VerifyToken(db *gorm.DB, token jwt.Token) (bool, error) {
	var sessions []Session
	result := db.Where("user_id = ?", user.ID).Find(&sessions)
	if result.Error != nil {
		return false, result.Error
	}
	found := false
	for _, session := range sessions {
		found = session.Token == token.Raw
		if found {
			break
		}
	}
	if !found {
		return false, nil
	}
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return false, err
	}
	sub, err := token.Claims.GetSubject()
	if err != nil {
		return false, err
	}
	return issuer == util.JWTIssuer && sub == user.Login, nil
}
