package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token  string `json:"token" gorm:"not null"`
	UserID uint   `json:"userID" gorm:"not null"`
}

// Simplified User object for use in session creation/verification.
type SessionUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
