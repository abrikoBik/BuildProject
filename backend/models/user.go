package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name" groups:"public" gorm:"unique;not null"`
	FullName string    `json:"fullName" groups:"public" gorm:"not null"`
	Email    string    `json:"email" groups:"public" gorm:"unique;not null"`
	Hash     []byte    `gorm:"not null"`
	Salt     []byte    `gorm:"not null"`
	Sessions []Session `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}
