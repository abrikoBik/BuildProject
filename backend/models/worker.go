package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	*User
	Description string `json:"description" groups:"public"`
}

type WorkerResponse struct {
	*Worker
}
