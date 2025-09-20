package models

import (
	"errors"

	"gorm.io/gorm"
)

type EstimateGroup struct {
	gorm.Model
	Name      string  `json:"name" groups:"public" gorm:"not null"`
	ProjectID uint    `json:"projectID" groups:"public" gorm:"not null"`
	TotalCost float64 `json:"totalCost" groups:"public" gorm:"not null"`
	Estimates []Estimate
}

type EstimateGroupResponse struct {
	*EstimateGroup
	ID uint `json:"id" groups:"register"`
}

type EstimateGroupRequest struct {
	*EstimateGroup
}

func (group *EstimateGroup) Get(db *gorm.DB, id uint) error {
	return db.Preload("Estimates").Where("id = ?", id).Take(&group).Error
}

func (group *EstimateGroup) Create(db *gorm.DB, request EstimateGroupRequest, projectID uint) error {
	var project Project
	if err := project.Get(db, request.ProjectID, []string{"EstimateGroups"}); err != nil {
		return err
	}
	group.Name = request.Name
	group.ProjectID = project.ID
	project.EstimateGroups = append(project.EstimateGroups, *group)
	return errors.Join(db.Save(&project).Error, db.Create(&group).Error)
}

func (group *EstimateGroup) Delete(db *gorm.DB) error {
	for _, estimate := range group.Estimates {
		if err := estimate.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&group).Error
}

func (group *EstimateGroup) CalculateCost(db *gorm.DB) error {
	if len(group.Estimates) == 0 {
		return nil
	}
	group.TotalCost = 0
	for _, estimate := range group.Estimates {
		group.TotalCost += estimate.Cost * float64(estimate.Count)
	}
	return db.Save(&group).Error
}
