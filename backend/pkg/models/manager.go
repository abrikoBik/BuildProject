package models

import (
	"errors"
	"mp1/pkg/util"
	"slices"

	"gorm.io/gorm"
)

type Manager struct {
	gorm.Model
	UserID  uint
	User    *User `gorm:"foreignKey"`
	Workers []Worker
}

type ManagerRequest struct {
	UserID uint `json:"userID"`
}

type ManagerResponse struct {
	UserID       uint   `json:"userID" groups:"public"`
	RegisteredAt string `json:"registeredAt" groups:"public"`
	ID           uint   `json:"id" groups:"register"`
}

func (manager *Manager) Get(db *gorm.DB, id uint) error {
	return db.Preload("Workers").Take(&manager).Error
}

func (manager *Manager) Create(db *gorm.DB, request ManagerRequest) error {
	var user User
	if err := user.Get(db, request.UserID); err != nil {
		return err
	}
	manager.User = &user
	manager.UserID = user.ID
	manager.Workers = make([]Worker, 0)
	return errors.Join(db.Create(&manager).Error, db.Save(&user).Error)
}

func (manager *Manager) AddWorker(db *gorm.DB, worker Worker) error {
	manager.Workers = append(manager.Workers, worker)
	return db.Save(&manager).Error
}

func (manager *Manager) DeleteWorker(db *gorm.DB, index int) error {
	if index < 0 || index >= len(manager.Workers) {
		return util.OutOfBoundsError
	}
	manager.Workers = slices.Delete(manager.Workers, index, index)
	return db.Save(&manager).Error
}
