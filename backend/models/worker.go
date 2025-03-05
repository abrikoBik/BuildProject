package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	User
	Description string `json:"description"`
}

type WorkerJSON struct {
	UserJSON
	Description string `json:"description"`
}
