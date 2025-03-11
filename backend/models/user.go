package models

import "gorm.io/gorm"

type User struct {
	Name        string `json:"name" gorm:"unique;not null"`
	FullName    string `json:"fullName" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	PhoneNumber uint   `json:"phoneNumber" gorm:"unique"`
}

type UserDB struct {
	gorm.Model
	User
	Hash     []byte `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
	Sessions []Session
}

type UserJSONIn struct {
	User     `json:"user"`
	Password string `json:"password"`
}

type UserJSONOut struct {
	User `json:"user"`
	ID   uint `json:"id"`
}
