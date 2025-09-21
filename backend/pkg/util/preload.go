package util

import "gorm.io/gorm"

func Preload(db *gorm.DB, preloads ...string) *gorm.DB {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	return tx
}
