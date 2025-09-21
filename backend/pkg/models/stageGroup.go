package models

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type StageGroup struct {
	gorm.Model
	Name      string `json:"name" groups:"public" gorm:"not null"`
	ProjectID uint   `json:"projectID" groups:"public" gorm:"not null"`
	StartTime pq.NullTime
	EndTime   pq.NullTime
	Stages    []Stage
}

type StageGroupRequest struct {
	*StageGroup
}

type StageGroupResponse struct {
	*StageGroup
	StartTime int64 `json:"startTime" groups:"public"`
	EndTime   int64 `json:"endTime" groups:"public"`
	CreatedAt int64 `json:"createdAt" groups:"public"`
}

func (group *StageGroup) Get(db *gorm.DB, id uint) error {
	return db.Preload("Stage").Where("id = ?", id).Take(&group).Error
}

func (group *StageGroup) Create(db *gorm.DB, request StageGroupRequest) error {
	var project Project
	if err := project.Get(db, request.ProjectID, []string{"StageGroups"}); err != nil {
		return err
	}
	group.Name = request.Name
	group.ProjectID = project.ID
	project.StageGroups = append(project.StageGroups, *group)
	return errors.Join(db.Save(&project).Error, db.Create(&group).Error)
}

func (group *StageGroup) Delete(db *gorm.DB) error {
	for _, stage := range group.Stages {
		if err := stage.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&group).Error
}

func (group *StageGroup) CalculateTimes(db *gorm.DB) error {
	if len(group.Stages) == 0 {
		return nil
	}
	group.StartTime = group.Stages[0].StartTime
	group.EndTime = group.Stages[len(group.Stages)-1].EndTime
	return db.Save(&group).Error
}
