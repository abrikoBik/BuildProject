package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token  string `json:"token" gorm:"not null"`
	UserID uint   `json:"userID" gorm:"not null"`
}

type SessionJSONIn struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SessionJSONOut struct {
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
}
