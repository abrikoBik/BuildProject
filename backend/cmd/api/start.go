package api

import (
	"errors"
	"mp1/pkg/api"
	"mp1/pkg/util"
	"net"
	"net/http"
)

func Start(hostname string, port uint, callback func()) {
	server = api.Init(hostname, port)
	util.Info("api", "starting @ %s", server.Addr)
	listener, err := net.Listen("tcp", server.Addr)
	util.Check(err)
	util.Info("api", "started")
	callback()
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		util.Error("api", err)
	}
}
