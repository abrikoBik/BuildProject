package db

import "mp1/pkg/util"

func Stop(callback func()) {
	dbLogger.Info("stopping")
	util.Check(sqlDatabase.Close())
	dbLogger.Info("stopped")
	callback()
}
