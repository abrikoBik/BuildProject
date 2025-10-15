package models

import (
	"mp1/pkg/util"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name         string `json:"name" groups:"public,list" gorm:"not null"`
	Status       bool   `json:"status" groups:"public,list" gorm:"not null;default:false"`
	StageGroupID uint   `gorm:"not null"`
	StartTime    pq.NullTime
	EndTime      pq.NullTime
}

type StageRequest struct {
	*Stage
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type StageResponse struct {
	*Stage
	StartTime int64 `json:"startTime" groups:"public,list"`
	EndTime   int64 `json:"endTime" groups:"public,list"`
	ID        uint  `json:"id" groups:"register,list"`
}

func (stage *Stage) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&stage).Error
}

func (stage *Stage) Create(db *gorm.DB, request StageRequest) error {
	stage.StageGroupID = request.StageGroupID
	stage.Name = request.Name
	stage.StartTime = util.UnmarshalTime(request.StartTime)
	stage.EndTime = util.UnmarshalTime(request.EndTime)
	return db.Create(&stage).Error
}

func (stage *Stage) Delete(db *gorm.DB) error {
	return db.Delete(&stage).Error
}

func (stage *Stage) Patch(db *gorm.DB, request StageRequest) error {
	if len(request.Name) > 0 {
		stage.Name = request.Name
	}
	if request.Status != stage.Status {
		stage.Status = request.Status
	}
	if request.StartTime != 0 {
		stage.StartTime = util.UnmarshalTime(request.StartTime)
	}
	if request.EndTime != 0 {
		stage.EndTime = util.UnmarshalTime(request.EndTime)
	}
	return db.Save(&stage).Error
}

func (stage Stage) GetGroup(db *gorm.DB) (group *StageGroup, err error) {
	err = group.Get(db, stage.StageGroupID)
	return
}
