package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	FullName    string `json:"fullName" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	PhoneNumber uint   `json:"phoneNumber" gorm:"unique"`
	Description string `json:"description"`
}
