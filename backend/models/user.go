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

type UserJSON struct {
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber uint   `json:"phoneNumber"`
}

type UserJSONIn struct {
	UserJSON `json:"user"`
	Password string `json:"password"`
}

type UserJSONOut struct {
	UserJSON `json:"user"`
	ID       uint `json:"id"`
}
