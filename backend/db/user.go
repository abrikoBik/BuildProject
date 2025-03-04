package db

import (
	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Users []User `json:"users"`
}

type Task struct {
	gorm.Model
	Start uint `json:"start" gorm:"not null"`
	End   uint `json:"end" gorm:"not null"`
}

type User struct {
	gorm.Model
	// TODO: decide whether name should be a phone number or an email address
	Name     string `json:"name" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Hash     []byte `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;not null"`

	// Only used in a JSON object.
	Password string `json:"password" gorm:"-"`
}

// User object for /userInfo endpoint.
type UserOut struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	ID       uint   `json:"id"`
}
