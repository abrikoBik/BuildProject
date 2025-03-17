package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string   `json:"name" groups:"public" gorm:"not null"`
	Users   []User   `json:"users" groups:"public" gorm:"many2many:user_projects;"`
	Workers []Worker `json:"workers" groups:"public" gorm:"many2many:worker_projects;"`
	Stages  []Stage  `json:"stages" groups:"public" gorm:"many2many:project_stages;"`
}

type ProjectResponse struct {
	*Project
}
