package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name      string      `json:"name" gorm:"not null"`
	StartTime pq.NullTime `json:"startTime"`
	EndTime   pq.NullTime `json:"endTime"`
}

type Project struct {
	gorm.Model
	Name    string   `json:"name" gorm:"not null"`
	Users   []User   `json:"users" gorm:"many2many:user_projects;"`
	Workers []Worker `json:"workers" gorm:"many2many:worker_projects;"`
	Stages  []Stage  `json:"stages" gorm:"many2many:project_stages;"`
}

type ProjectResponse struct {
	*Project
}

type StageResponse struct {
	*Stage
	ProjectID uint `json:"project"`
}
