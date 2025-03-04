package models

import (
	"time"

	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}

type ProjectStage struct {
	gorm.Model
	Project   Project   `json:"project"`
	Stages    []Stage   `json:"stages" gorm:"many2many:project_stages;"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type Project struct {
	gorm.Model
	Name    string   `json:"name" gorm:"not null"`
	Users   []User   `json:"users" gorm:"many2many:user_projects;"`
	Workers []Worker `json:"workers" gorm:"many2many:worker_projects;"`
}
