package api

import "mp1/pkg/util"

func Stop(callback func()) {
	util.Info("api", "stopping")
	util.Check(server.Close())
	util.Info("api", "stopped")
	callback()
}
