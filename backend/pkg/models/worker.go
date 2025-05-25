package models

import (
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Role   string `json:"role" groups:"public" gorm:"not null"`
	UserID uint   `json:"userID" groups:"public" gorm:"not null"`
}

type WorkerRequest struct {
	*Worker
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
	worker.Role = request.Role
	var user User
	if err := db.Where("id = ?", request.UserID).Take(&user).Error; err != nil {
		return err
	}
	worker.UserID = request.UserID
	return db.Create(&worker).Error
}

func (worker *Worker) Patch(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) > 0 {
		worker.Role = request.Role
	}
	return db.Save(&worker).Error
}
