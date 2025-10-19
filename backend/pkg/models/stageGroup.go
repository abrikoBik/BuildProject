package models

import (
	"log"
	"mp1/pkg/util"
	"sort"
	"time"

	"gorm.io/gorm"
)

type StageGroup struct {
	gorm.Model
	Name      string  `json:"name" groups:"public,list" gorm:"not null"`
	Project   Project `gorm:"foreignKey:ProjectID;not null"`
	ProjectID uint    `gorm:"not null"`
	Stages    []Stage
}

type StageGroupRequest struct {
	*StageGroup
}

type StageGroupResponse struct {
	*StageGroup
	StartTime       int64 `json:"startTime" groups:"public"`
	EndTime         int64 `json:"endTime" groups:"public"`
	TotalStages     int   `json:"totalStages" groups:"public"`
	CompletedStages int   `json:"completedStages" groups:"public"`
	ID              uint  `json:"id" groups:"register,list"`
}

func (group *StageGroup) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&group).Error
}

func (group *StageGroup) Create(db *gorm.DB, request StageGroupRequest) error {
	var project Project
	if err := project.Get(db, request.ProjectID, "StageGroups"); err != nil {
		return err
	}
	group.Name = request.Name
	group.ProjectID = project.ID
	return db.Create(&group).Error
}

func (group *StageGroup) Patch(db *gorm.DB, request StageGroupRequest) error {
	if len(request.Name) > 0 {
		group.Name = request.Name
	}
	return db.Save(&group).Error
}

func (group *StageGroup) Delete(db *gorm.DB) error {
	for _, stage := range group.Stages {
		if err := stage.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&group).Error
}

func (group StageGroup) CalculateTimes(db *gorm.DB) (start time.Time, end time.Time) {
	if len(group.Stages) == 0 {
		return
	}
	sort.Slice(group.Stages, func(i, j int) bool {
		return group.Stages[i].StartTime.Time.UnixMilli() < group.Stages[j].StartTime.Time.UnixMilli()
	})
	log.Println(group.Stages[0])
	start = group.Stages[0].StartTime.Time
	sort.Slice(group.Stages, func(i, j int) bool {
		return group.Stages[i].EndTime.Time.UnixMilli() > group.Stages[j].EndTime.Time.UnixMilli()
	})
	log.Println(group.Stages[0])
	end = group.Stages[0].EndTime.Time
	return
}

func (group StageGroup) CalculateStages(db *gorm.DB) (total, done int) {
	total = len(group.Stages)
	for _, s := range group.Stages {
		if s.Status {
			done += 1
		}
	}
	return
}
