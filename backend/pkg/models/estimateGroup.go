package models

import (
	"errors"
	"mp1/pkg/l10n"
	"mp1/pkg/util"
	"slices"

	"gorm.io/gorm"
)

type EstimateGroup struct {
	gorm.Model
	Name      string `json:"name" groups:"public,list" gorm:"not null"`
	ProjectID uint   `gorm:"not null"`
	Estimates []Estimate
}

type EstimateGroupRequest struct {
	*EstimateGroup
}

type EstimateGroupResponse struct {
	*EstimateGroup
	TotalCost float64 `json:"totalCost" groups:"public,list"`
	ID        uint    `json:"id" groups:"register,list"`
}

func (group *EstimateGroup) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&group).Error
}

func (group *EstimateGroup) Create(db *gorm.DB, request EstimateGroupRequest) error {
	var project Project
	if err := project.Get(db, request.ProjectID, "EstimateGroups.Estimates"); err != nil {
		return err
	}
	group.Name = request.Name
	group.ProjectID = project.ID
	return errors.Join(db.Create(&group).Error, project.AddEstimateGroup(db, *group))
}

func (group *EstimateGroup) Patch(db *gorm.DB, request EstimateGroupRequest) error {
	if len(request.Name) > 0 {
		group.Name = request.Name
	}
	return db.Save(&group).Error
}

func (group *EstimateGroup) Delete(db *gorm.DB) error {
	for _, estimate := range group.Estimates {
		if err := estimate.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&group).Error
}

func (group *EstimateGroup) AddEstimate(db *gorm.DB, estimate Estimate) error {
	group.Estimates = append(group.Estimates, estimate)
	return db.Save(&group).Error
}

func (group *EstimateGroup) RemoveEstimate(db *gorm.DB, index int) error {
	if index < 0 || index >= len(group.Estimates) {
		return l10n.OutOfBoundsError
	}
	group.Estimates = slices.Delete(group.Estimates, index, index+1)
	return db.Save(&group).Error
}

func (group *EstimateGroup) CalculateCost() (totalCost float64) {
	if len(group.Estimates) == 0 {
		return
	}
	for _, estimate := range group.Estimates {
		totalCost += estimate.Cost * float64(estimate.Count)
	}
	return
}
