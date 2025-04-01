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
	Password string `json:"password"`
}

type WorkerResponse struct {
	*Worker
	RegisteredAt string `json:"registeredAt" groups:"public"`
	ID           uint   `json:"id" groups:"register"`
}

func (worker *Worker) Create(db *gorm.DB, request WorkerRequest) error {
	var user User
	if err := user.FindByID(db, worker.UserID); err != nil {
		return err
	}
	match, err := user.VerifyPassword(db, request.Password)
	if err != nil {
		return err
	}
	if !match {
		return util.WrongPasswordError
	}
	worker.Role = request.Role
	return nil
}
