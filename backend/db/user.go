package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// TODO: decide whether name should be a phone number or an email address
	Name     string `json:"name" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Hash     string `json:"hash" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;not null"`
}
