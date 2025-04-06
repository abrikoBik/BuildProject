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
	if len(request.Role) == 0 {
		return util.InvalidRoleError
	}
	worker.Role = request.Role
	if len(request.Password) < util.MinPasswordLength {
		return util.InsecurePasswordError
	}
	var user User
	if err := db.Where("id = ?", request.UserID).Take(&user).Error; err != nil {
		return err
	}
	println("user found")
	match, err := user.VerifyPassword(db, request.Password)
	if err != nil {
		return err
	}
	if !match {
		return util.WrongPasswordError
	}
	worker.UserID = request.UserID
	return nil
}
