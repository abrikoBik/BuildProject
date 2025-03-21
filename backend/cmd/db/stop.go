package db

import "mp1/pkg/util"

func Stop(callback func()) {
	util.Info("db", "stopping")
	util.Check(sqlDatabase.Close())
	util.Info("db", "stopped")
	callback()
}
