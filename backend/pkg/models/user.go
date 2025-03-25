package models

import (
	"bytes"
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
	Worker   Worker    `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

type UserResponse struct {
	*User
	ID           uint   `json:"id" groups:"register"`
	RegisteredAt string `json:"registeredAt" groups:"public"`
}

func (user *User) Create(db *gorm.DB, request UserRequest) (err error) {
	if user.Salt, err = util.RandomBytes(16); err != nil {
		return err
	}
	if user.Hash, err = util.GenerateHash([]byte(request.Password), user.Salt); err != nil {
		return err
	}
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) FindByID(db *gorm.DB, id uint) error {
	result := db.Where("id = ?", id).Take(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) FindByLogin(db *gorm.DB, login string) error {
	result := db.Where("login = ?", login).Take(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user User) VerifyPassword(db *gorm.DB, password string) (bool, error) {
	hash, err := util.GenerateHash([]byte(password), user.Salt)
	if err != nil {
		return false, err
	}
	return (bytes.Compare(user.Hash, hash) == 0), nil
}
