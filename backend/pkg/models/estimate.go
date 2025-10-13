package models

import (
	"gorm.io/gorm"
)

type Estimate struct {
	gorm.Model
	Name            string  `json:"name" groups:"public,list" gorm:"not null"`
	Cost            float64 `json:"cost" groups:"public,list" gorm:"not null"`
	Count           uint    `json:"count" groups:"public,list" gorm:"not null"`
	Unit            string  `json:"unit" groups:"public,list" gorm:"not null"`
	EstimateGroupID uint    `gorm:"not null"`
}

type EstimateRequest struct {
	*Estimate
}

type EstimateResponse struct {
	*Estimate
	ID uint `json:"id" groups:"register,list"`
}

func (estimate *Estimate) Get(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Take(&estimate).Error
}

func (estimate *Estimate) Create(db *gorm.DB, request EstimateRequest) error {
	var estimateGroup EstimateGroup
	if err := estimateGroup.Get(db, request.EstimateGroupID, "Estimates"); err != nil {
		return err
	}
	estimate.EstimateGroupID = request.EstimateGroupID
	estimate.Name = request.Name
	estimate.Cost = request.Cost
	estimate.Count = request.Count
	estimate.Unit = request.Unit
	return estimateGroup.AddEstimate(db, *estimate)
}

func (estimate *Estimate) Patch(db *gorm.DB, request EstimateRequest) error {
	if len(request.Name) > 0 {
		estimate.Name = request.Name
	}
	if request.Cost != 0 {
		estimate.Cost = request.Cost
	}
	if request.Count != 0 {
		estimate.Count = request.Count
	}
	if len(request.Unit) > 0 {
		estimate.Unit = request.Unit
	}
	return db.Save(&estimate).Error
}

// TODO: remove? EstimateGroup.RemoveEstimate might supersede this
func (estimate *Estimate) Delete(db *gorm.DB) error {
	return db.Delete(&estimate).Error
}
