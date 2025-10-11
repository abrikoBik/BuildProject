package models

import (
	"log"
	"mp1/pkg/util"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name         string `json:"name" groups:"public,list" gorm:"not null"`
	Description  string `json:"description" groups:"public"`
	Status       bool   `json:"status" groups:"public,list" gorm:"not null;default:false"`
	StageGroupID uint   `json:"groupID" gorm:"not null"`
	StartTime    pq.NullTime
	EndTime      pq.NullTime
}

type StageRequest struct {
	*Stage
	StartTime int64 `json:"startTime" groups:"public"`
	EndTime   int64 `json:"endTime" groups:"public"`
}

type StageResponse struct {
	*Stage
	StartTime int64 `json:"startTime" groups:"public"`
	EndTime   int64 `json:"endTime" groups:"public"`
	ID        uint  `json:"id" groups:"register,list"`
}

func (stage *Stage) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&stage).Error
}

func (stage *Stage) Create(db *gorm.DB, request StageRequest) error {
	stage.StageGroupID = request.StageGroupID
	// group, err := stage.GetGroup(db)
	// if err != nil {
	// 	return err
	// }
	stage.Name = request.Name
	stage.Description = request.Description
	stage.StartTime = util.UnmarshalTime(request.StartTime)
	stage.EndTime = util.UnmarshalTime(request.EndTime)
	log.Println(stage, request)
	// return errors.Join(db.Create(&stage).Error, group.CalculateTimes(db), db.Save(&group).Error)
	return db.Create(&stage).Error
}

func (stage *Stage) Delete(db *gorm.DB) error {
	return db.Delete(&stage).Error
}

func (stage *Stage) Patch(db *gorm.DB, request StageRequest) error {
	if len(request.Name) > 0 {
		stage.Name = request.Name
	}
	if len(request.Description) > 0 {
		stage.Description = request.Description
	}
	if request.StartTime != 0 {
		stage.StartTime = util.UnmarshalTime(request.StartTime)
	}
	if request.EndTime != 0 {
		stage.EndTime = util.UnmarshalTime(request.EndTime)
	}
	return db.Save(&stage).Error
}

func (stage *Stage) Finish(db *gorm.DB) error {
	stage.Status = true
	return db.Save(&stage).Error
}

func (stage Stage) GetGroup(db *gorm.DB) (group *StageGroup, err error) {
	err = group.Get(db, stage.StageGroupID)
	return
}
