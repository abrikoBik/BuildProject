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
	apiLogger.Info("starting @ %s", server.Addr)
	listener, err := net.Listen("tcp", server.Addr)
	util.Check(err)
	apiLogger.Info("started")
	callback()
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		apiLogger.Error(err)
	}
}
