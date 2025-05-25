package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string  `json:"name" groups:"public" gorm:"not null"`
	Description string  `json:"description" groups:"public"`
	User        User    `json:"user" groups:"public" gorm:"many2many:user_projects"`
	Stages      []Stage `json:"stages" groups:"public"`
}

type ProjectResponse struct {
	*Project
	CreatedAt string `json:"createdAt" groups:"public"`
	ID        uint   `json:"id" groups:"register"`
}
