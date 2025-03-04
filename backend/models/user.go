package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Hash     []byte `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;not null"`

	// Only used in a JSON object.
	Password string `json:"password" gorm:"-"`
}

type UserOut struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	ID       uint   `json:"id"`
}
