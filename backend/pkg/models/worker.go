package models

import (
	"errors"
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Role      string `json:"role" groups:"public" gorm:"not null"`
	ManagerID uint
	Manager   *Manager `gorm:"foreignKey:ManagerID"`
}

type WorkerRequest struct {
	*Worker
	ManagerID uint `json:"managerID"`
}

type WorkerResponse struct {
	*Worker
	RegisteredAt string `json:"registeredAt" groups:"public"`
	ID           uint   `json:"id" groups:"register"`
}

func (worker *Worker) Get(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Take(&worker).Error
}

func (worker *Worker) Create(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) == 0 {
		return util.InvalidRoleError
	}
	var manager Manager
	if err := manager.Get(db, request.ManagerID); err != nil {
		return err
	}
	worker.Role = request.Role
	worker.Manager = &manager
	worker.ManagerID = manager.ID
	if manager.Worker != nil {
		return util.WorkerExistsError
	}
	return errors.Join(db.Create(&worker).Error, db.Save(&manager).Error)
}

func (worker *Worker) Patch(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) > 0 {
		worker.Role = request.Role
	}
	return db.Save(&worker).Error
}
