package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// TODO: decide whether name should be a phone number or an email address
	Name     string `json:"name" gorm:"unique;not null"`
	FullName string `json:"fullName" gorm:"not null"`
	Hash     string `json:"hash" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;not null"`
}

// Simplified User object for use in session creation/verification.
type SessionUser struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
	Salt string `json:"salt"`
}

type Session struct {
	gorm.Model
	Name      string    `json:"name" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
}
