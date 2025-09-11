package models

import (
	"errors"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name      string `json:"name" groups:"public" gorm:"not null"`
	ProjectID uint   `json:"projectID" groups:"public" gorm:"not null"`
	StartTime pq.NullTime
	EndTime   pq.NullTime
	Stages    []Stage
}

type GroupRequest struct {
	*Group
}

type GroupResponse struct {
	*Group
	StartTime int64 `json:"startTime" groups:"public"`
	EndTime   int64 `json:"endTime" groups:"public"`
	CreatedAt int64 `json:"createdAt" groups:"public"`
}

func (group *Group) Get(db *gorm.DB, id uint) error {
	return db.Preload("Stages").Where("id = ?", id).Take(&group).Error
}

func (group *Group) Create(db *gorm.DB, request GroupRequest) error {
	var project Project
	if err := project.Get(db, request.ProjectID, []string{"Groups"}); err != nil {
		return err
	}
	group.Name = request.Name
	group.ProjectID = project.ID
	project.Groups = append(project.Groups, *group)
	return errors.Join(db.Save(&project).Error, db.Create(&group).Error)
}

func (group *Group) Delete(db *gorm.DB) error {
	for _, stage := range group.Stages {
		if err := stage.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&group).Error
}

func (group *Group) CalculateTimes(db *gorm.DB) error {
	if len(group.Stages) == 0 {
		return nil
	}
	group.StartTime = group.Stages[0].StartTime
	group.EndTime = group.Stages[len(group.Stages)-1].EndTime
	return db.Save(&group).Error
}
