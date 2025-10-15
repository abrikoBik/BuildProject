package models

import (
	"mp1/pkg/util"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type StageGroup struct {
	gorm.Model
	Name      string      `json:"name" groups:"public,list" gorm:"not null"`
	ProjectID uint        `gorm:"not null"`
	StartTime pq.NullTime `gorm:"not null"`
	EndTime   pq.NullTime `gorm:"not null"`
	Stages    []Stage
}

type StageGroupRequest struct {
	*StageGroup
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type StageGroupResponse struct {
	*StageGroup
	StartTime int64 `json:"startTime" groups:"public"`
	EndTime   int64 `json:"endTime" groups:"public"`
	ID        uint  `json:"id" groups:"register,list"`
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
	group.StartTime = util.UnmarshalTime(request.StartTime)
	group.EndTime = util.UnmarshalTime(request.EndTime)
	group.ProjectID = project.ID
	return db.Create(&group).Error
}

func (group *StageGroup) Patch(db *gorm.DB, request StageGroupRequest) error {
	if len(request.Name) > 0 {
		group.Name = request.Name
	}
	if request.StartTime != 0 {
		group.StartTime = util.UnmarshalTime(request.StartTime)
	}
	if request.EndTime != 0 {
		group.EndTime = util.UnmarshalTime(request.EndTime)
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

// TODO: rework
// func (group *StageGroup) CalculateTimes(db *gorm.DB) error {
// 	if len(group.Stages) == 0 {
// 		return nil
// 	}
// 	sort.Slice(group.Stages, func(i, j int) bool {
// 		return group.Stages[i].StartTime > group.Stages[j].StartTime
// 	})
// 	group.StartTime = group.Stages[0].StartTime
// 	group.EndTime = group.Stages[len(group.Stages)-1].EndTime
// 	return db.Save(&group).Error
// }
