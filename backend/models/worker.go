package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	*User
	Description string
}

type WorkerResponse struct {
	*UserResponse
	Description string `json:"description"`
}
