package models

import (
	"errors"
	"mp1/pkg/l10n"
	"mp1/pkg/util"
	"slices"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name         string `json:"name" groups:"public,list" gorm:"not null"`
	Description  string `json:"description" groups:"public"`
	Status       bool   `json:"status" groups:"public,list" gorm:"not null;default:false"`
	Index        int    `json:"index" groups:"public,list" gorm:"not null"`
	StageGroupID uint   `json:"groupID" gorm:"not null"`
	StartTime    pq.NullTime
	EndTime      pq.NullTime
	Workers      []Worker `gorm:"many2many:stage_workers"`
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
	CreatedAt int64 `json:"createdAt" groups:"public"`
	ID        uint  `json:"id" groups:"register,list"`
}

func (stage *Stage) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&stage).Error
}

func (stage *Stage) Create(db *gorm.DB, request StageRequest) error {
	group, err := stage.GetGroup(db)
	if err != nil {
		return err
	}
	if request.Index < 0 || request.Index >= len(group.Stages) {
		return l10n.OutOfBoundsError
	}
	stage.Name = request.Name
	stage.Description = request.Description
	stage.StartTime = util.UnmarshalTime(request.StartTime)
	stage.EndTime = util.UnmarshalTime(request.EndTime)
	stage.StageGroupID = group.ID
	stage.Index = request.Index
	stage.Workers = make([]Worker, 0)
	group.AddStage(db, *stage)
	return errors.Join(db.Create(&stage).Error, group.CalculateTimes(db), db.Save(&group).Error)
}

func (stage *Stage) Delete(db *gorm.DB) error {
	group, err := stage.GetGroup(db)
	if err != nil {
		return err
	}
	if stage.Index < 0 || stage.Index >= len(group.Stages) {
		return l10n.OutOfBoundsError
	}
	group.Stages = slices.Delete(group.Stages, stage.Index, stage.Index+1)
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
		return l10n.OutOfBoundsError
	}
	stage.Workers = slices.Delete(stage.Workers, index, index+1)
	return db.Save(&stage).Error
}

func (stage Stage) GetGroup(db *gorm.DB) (group *StageGroup, err error) {
	err = group.Get(db, stage.StageGroupID)
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
