package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token  string `json:"token" groups:"public" gorm:"not null"`
	UserID uint   `json:"userID" groups:"public" gorm:"not null"`
}

type SessionRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SessionResponse struct {
	*Session
}
