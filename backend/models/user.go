package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `gorm:"unique;not null"`
	FullName    string    `gorm:"not null"`
	Email       string    `gorm:"not null"`
	PhoneNumber uint      `gorm:"unique"`
	Hash        []byte    `gorm:"not null"`
	Salt        []byte    `gorm:"not null"`
	Sessions    []Session `gorm:"foreignKey:UserID"`
}

type UserJSONIn struct {
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber uint   `json:"phoneNumber"`
	Password    string `json:"password"`
}

type UserJSONOut struct {
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber uint   `json:"phoneNumber"`
	ID          uint   `json:"id"`
}
