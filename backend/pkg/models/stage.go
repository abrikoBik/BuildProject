package models

import (
	"mp1/pkg/util"
	"sync"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Stage struct {
	gorm.Model
	Name        string      `json:"name" groups:"public" gorm:"not null"`
	Description string      `json:"description" groups:"public"`
	StartTime   pq.NullTime `json:"startTime" groups:"public" gorm:"not null"`
	EndTime     pq.NullTime `json:"endTime" groups:"public" gorm:"not null"`
	ProjectID   uint        `json:"projectID" groups:"public"`
	Workers     []uint      `json:"workers" groups:"public"`
	Status      bool        `json:"status" groups:"public" gorm:"not null;default:false"`
}

type StageRequest struct {
	*Stage
}

type StageResponse struct {
	*Stage
	ID uint `json:"id" groups:"register"`
}

func (stage *Stage) Get(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Take(&stage).Error
}

func (stage *Stage) Create(db *gorm.DB, request StageRequest) error {
	if !stage.VerifyWorkers(db, request.Workers) {
		return util.WorkerNotFoundError
	}
	stage.Name = request.Stage.Name
	stage.Description = request.Stage.Description
	stage.Workers = request.Stage.Workers
	stage.Status = request.Stage.Status
	// TODO: add ProjectID checks
	stage.ProjectID = request.Stage.ProjectID
	return db.Create(&stage).Error
}

func (stage *Stage) Patch(db *gorm.DB, request StageRequest) error {
	if len(request.Stage.Name) > 0 {
		stage.Name = request.Stage.Name
	}
	if len(request.Stage.Description) > 0 {
		stage.Description = request.Stage.Description
	}
	if len(request.Stage.Name) > 0 {
		stage.Name = request.Stage.Name
	}
	return db.Save(&stage).Error
}

func (stage *Stage) Finish(db *gorm.DB) error {
	stage.Status = true
	return db.Save(&stage).Error
}

func (stage Stage) VerifyWorkers(db *gorm.DB, ids []uint) (result bool) {
	var (
		wg      sync.WaitGroup
		errored = make(chan bool)
	)
	wg.Add(len(ids))
	for _, workerID := range ids {
		go func(erroredChan chan<- bool, current uint) {
			defer wg.Done()
			var worker Worker
			if worker.Get(db, current) != nil {
				erroredChan <- true
				close(erroredChan)
			}
		}(errored, workerID)
	}
	go func(erroredChan chan<- bool) {
		wg.Wait()
		if result {
			return
		}
		erroredChan <- false
		close(erroredChan)
	}(errored)
	result = <-errored
	return
}
