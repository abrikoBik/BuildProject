package models

import (
	"mp1/pkg/util"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EstimateGroup struct {
	gorm.Model
	Name      string  `json:"name" groups:"public,list" gorm:"not null"`
	Project   Project `gorm:"foreignKey:ProjectID;not null"`
	ProjectID uint    `gorm:"not null"`
	Estimates []Estimate
}

type EstimateGroupRequest struct {
	*EstimateGroup
}

type EstimateGroupResponse struct {
	*EstimateGroup
	TotalCost float64 `json:"totalCost" groups:"public"`
	ID        uint    `json:"id" groups:"register,list"`
}

func (group *EstimateGroup) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&group).Error
}

func (group *EstimateGroup) Create(db *gorm.DB, request EstimateGroupRequest) error {
	group.ProjectID = request.ProjectID
	group.Name = request.Name
	return db.Create(&group).Error
}

func (group *EstimateGroup) Patch(db *gorm.DB, request EstimateGroupRequest) error {
	if len(request.Name) > 0 {
		group.Name = request.Name
	}
	return db.Save(&group).Error
}

func (group *EstimateGroup) Delete(db *gorm.DB) error {
	return db.Delete(&group).Error
}

func (group *EstimateGroup) AfterDelete(db *gorm.DB) error {
	return db.Clauses(clause.Returning{}).Where("estimate_group_id = ?", group.ID).Delete(&Estimate{}).Error
}

func (group *EstimateGroup) CalculateCost() (totalCost float64) {
	for _, estimate := range group.Estimates {
		totalCost += estimate.Cost * float64(estimate.Count)
	}
	return
}
