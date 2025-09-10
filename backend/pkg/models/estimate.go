package models

import "gorm.io/gorm"

type Estimate struct {
	gorm.Model
	Name  string  `json:"name" groups:"public" gorm:"not null"`
	Cost  float32 `json:"cost" groups:"public" gorm:"not null"`
	Count uint    `json:"count" groups:"public" gorm:"not null"`
}

type EstimateResponse struct {
	*Estimate
	ID uint `json:"id" groups:"register"`
}

type EstimateRequest struct {
	*Estimate
}

func (estimate *Estimate) Get(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Take(&estimate).Error
}

func (estimate *Estimate) Create(db *gorm.DB, request EstimateRequest) error {
	estimate.Name = request.Name
	estimate.Cost = request.Cost
	estimate.Count = request.Count
	return db.Create(&estimate).Error
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
	return db.Save(&estimate).Error
}

func (estimate *Estimate) Delete(db *gorm.DB) error {
	return db.Delete(&estimate).Error
}
