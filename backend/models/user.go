package models

import (
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique;not null"`
	FullName    string    `json:"fullName" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null"`
	PhoneNumber uint      `json:"phoneNumber" gorm:"unique"`
	Hash        []byte    `gorm:"not null"`
	Salt        []byte    `gorm:"not null"`
	Sessions    []Session `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

func (user *UserRequest) Bind(r *http.Request) error {
	nonNullable := []any{user.Name, user.FullName, user.Email, user.PhoneNumber}
	for _, field := range nonNullable {
		if field == nil {
			return errors.New(fmt.Sprintf("missing field %s for UserRequest\n", field))
		}
	}
	return nil
}

type UserResponse struct {
	*User
}
