package models

import "gorm.io/gorm"

type WorkerDB struct {
	gorm.Model
	User
	Description string
}

type WorkerJSON struct {
	UserJSONOut
	Description string `json:"description"`
}
