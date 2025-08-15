package models

import (
	"errors"
	"mp1/pkg/util"
	"slices"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	StartTime   pq.NullTime
	EndTime     pq.NullTime
	Status      bool     `gorm:"not null;default:false"`
	GroupID     uint     `gorm:"not null"`
	Index       int      `gorm:"not null"`
	Workers     []Worker `gorm:"many2many:stage_workers"`
}

type StageRequest struct {
	Name        string `json:"name" groups:"public"`
	Description string `json:"description" groups:"public"`
	StartTime   int64  `json:"startTime" groups:"public"`
	EndTime     int64  `json:"endTime" groups:"public"`
	GroupID     uint   `json:"projectID" groups:"public"`
	Index       int    `json:"index" groups:"public"`
}

type StageResponse struct {
	Name        string `json:"name" groups:"public"`
	Description string `json:"description" groups:"public"`
	StartTime   int64  `json:"startTime" groups:"public"`
	EndTime     int64  `json:"endTime" groups:"public"`
	GroupID     uint   `json:"projectID" groups:"public"`
	Index       int    `json:"index" groups:"public"`
	CreatedAt   int64  `json:"createdAt" groups:"public"`
	ID          uint   `json:"id" groups:"register"`
}

func (stage *Stage) Get(db *gorm.DB, id uint) error {
	return db.Preload("Workers").Where("id = ?", id).Take(&stage).Error
}

func (stage *Stage) Create(db *gorm.DB, request StageRequest) error {
	group, err := stage.GetGroup(db)
	if err != nil {
		return err
	}
	if request.Index < 0 || request.Index >= len(group.Stages) {
		return util.OutOfBoundsError
	}
	stage.Name = request.Name
	stage.Description = request.Description
	stage.StartTime = util.UnmarshalTime(request.StartTime)
	stage.EndTime = util.UnmarshalTime(request.EndTime)
	stage.GroupID = group.ID
	stage.Index = request.Index
	stage.Workers = make([]Worker, 0)
	group.Stages = slices.Insert(group.Stages, stage.Index, *stage)
	return errors.Join(group.CalculateTimes(db), db.Create(&stage).Error)
}

func (stage *Stage) Delete(db *gorm.DB) error {
	group, err := stage.GetGroup(db)
	if err != nil {
		return err
	}
	if stage.Index < 0 || stage.Index >= len(group.Stages) {
		return util.OutOfBoundsError
	}
	group.Stages = slices.Delete(group.Stages, stage.Index, stage.Index)
	return errors.Join(group.CalculateTimes(db), db.Delete(&stage).Error)
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

func (stage *Stage) AddWorker(db *gorm.DB, worker Worker) error {
	stage.Workers = append(stage.Workers, worker)
	return db.Save(&stage).Error
}

func (stage *Stage) RemoveWorker(db *gorm.DB, index int) error {
	if index < 0 || index >= len(stage.Workers) {
		return util.OutOfBoundsError
	}
	stage.Workers = slices.Delete(stage.Workers, index, index)
	return db.Save(&stage).Error
}

func (stage Stage) GetGroup(db *gorm.DB) (group *Group, err error) {
	err = group.Get(db, stage.GroupID)
	return
}

// func (stage Stage) VerifyWorkers(db *gorm.DB, ids []uint) (result bool) {
// 	var (
// 		wg      sync.WaitGroup
// 		errored = make(chan bool)
// 	)
// 	wg.Add(len(ids))
// 	for _, workerID := range ids {
// 		go func(erroredChan chan<- bool, current uint) {
// 			defer wg.Done()
// 			var worker Worker
// 			if worker.Get(db, current) != nil {
// 				erroredChan <- true
// 				close(erroredChan)
// 			}
// 		}(errored, workerID)
// 	}
// 	go func(erroredChan chan<- bool) {
// 		wg.Wait()
// 		if result {
// 			return
// 		}
// 		erroredChan <- false
// 		close(erroredChan)
// 	}(errored)
// 	result = <-errored
// 	return
// }
