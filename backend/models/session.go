package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token string `json:"token" gorm:"not null"`
}

type SessionJSONIn struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SessionJSONOut struct {
	User  `json:"user"`
	Token string `json:"token"`
}
