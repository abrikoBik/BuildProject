package api

import "mp1/pkg/util"

func Stop(callback func()) {
	apiLogger.Info("api", "stopping")
	util.Check(server.Close())
	apiLogger.Info("api", "stopped")
	callback()
}
