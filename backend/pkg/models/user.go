package models

import (
	"mp1/pkg/util"

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
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

func (user *User) Create(db *gorm.DB, r UserRequest) (err error) {
	if user.Salt, err = util.RandomBytes(128); err != nil {
		return err
	}
	if user.Hash, err = util.GenerateHash([]byte(r.Password), user.Salt); err != nil {
		return err
	}
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
