package db

import (
	"mp1/pkg/db"
	"mp1/pkg/models"
	"mp1/pkg/util"
)

func Migrate(callback func()) {
	util.Info("db", "running migrations")
	toMigrate := map[string]any{
		"user":    models.User{},
		"session": models.Session{},
		"worker":  models.Worker{},
		"stage":   models.Stage{},
		"project": models.Project{},
	}
	for name, model := range toMigrate {
		util.Check(db.Database.Debug().AutoMigrate(&model))
		util.Debug("db", "model", name)
	}
	util.Info("db", "migrations complete")
	callback()
}
