package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"passwordHash"`
	Role         string `json:"role"`
}
