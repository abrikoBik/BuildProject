package db

import (
	"mp1/pkg/db"
	"mp1/pkg/models"
	"mp1/pkg/util"
)

func Migrate(callback func()) {
	migrationLogger := util.CreateLogger("db:Migrate")
	migrationLogger.Info("running migrations")
	toMigrate := map[string]any{
		"User":    models.User{},
		"Session": models.Session{},
		"Worker":  models.Worker{},

		"Project":       models.Project{},
		"Attachment":    models.Attachment{},
		"EstimateGroup": models.EstimateGroup{},
		"StageGroup":    models.StageGroup{},

		"Estimate": models.Estimate{},
		"Stage":    models.Stage{},
	}
	for name, model := range toMigrate {
		migrationLogger.Info("(%s) running migrations", name)
		util.Check(db.Database.Debug().AutoMigrate(&model))
		migrationLogger.Info("(%s) done", name)
	}
	migrationLogger.Info("migrations complete")
	callback()
}
