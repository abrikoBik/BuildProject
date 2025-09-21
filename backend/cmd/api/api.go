package api

import (
	"mp1/pkg/util"
	"net/http"
)

var (
	server    *http.Server
	apiLogger = util.CreateLogger("cmd/api")
)
