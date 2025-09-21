package db

import (
	"mp1/pkg/db"
	"mp1/pkg/util"
)

func Start(hostname string, port uint, name string, callback func()) {
	dbLogger.Info("connecting to %s @ %s:%d", name, hostname, port)
	util.Check(db.Open(hostname, port, name))
	var err error
	sqlDatabase, err = db.Database.DB()
	util.Check(err)
	dbLogger.Info("connected")
	callback()
}
