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
	Name    string     `json:"name" gorm:"not null"`
	Users   []User     `json:"users" gorm:"many2many:user_projects;"`
	Workers []WorkerDB `json:"workers" gorm:"many2many:worker_projects;"`
	Stages  []Stage    `json:"stages" gorm:"many2many:project_stages;"`
}

type ProjectJSON struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`

	// Array of IDs
	Users   []uint `json:"users"`
	Workers []uint `json:"workers"`
	Stages  []uint `json:"stages"`
}

type StageJSON struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`

	StartTime uint `json:"startTime"`
	EndTime   uint `json:"endTime"`

	// Project ID
	Project uint `json:"project"`
}
