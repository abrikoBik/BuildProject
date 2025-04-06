package models

import (
	"bytes"
	"mp1/pkg/util"
	"net/mail"

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
	err = db.Create(&user).Error
	return
}

func (user *User) Patch(db *gorm.DB, request UserRequest) (err error) {
	if len(request.Email) > 0 {
		if _, err = mail.ParseAddress(request.Email); err != nil {
			return util.InvalidEmailError
		}
		user.Email = request.Email
	}
	if len(request.Password) > 0 {
		if len(request.Password) < util.MinPasswordLength {
			return util.InsecurePasswordError
		}
		if err = user.PreparePassword(request.Password); err != nil {
			return
		}
	}
	if len(request.Name) > 0 {
		user.Name = cases.Title(language.Russian).String(request.Name)
	}
	err = db.Save(&user).Error
	return
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
