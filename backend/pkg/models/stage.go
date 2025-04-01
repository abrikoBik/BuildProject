package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name      string      `json:"name" groups:"public" gorm:"not null"`
	StartTime pq.NullTime `json:"startTime" groups:"public"`
	EndTime   pq.NullTime `json:"endTime" groups:"public"`
	Status    bool        `json:"status" groups:"public" gorm:"not null;default:false"`
}

type StageResponse struct {
	*Stage
	ProjectID uint `json:"project" groups:"public"`
}
