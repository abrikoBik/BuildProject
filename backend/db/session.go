package db

import "gorm.io/gorm"

// Simplified User object for use in session creation/verification.
type SessionUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Session struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Token string `json:"token" gorm:"not null"`
}
