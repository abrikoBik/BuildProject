package db

import (
	"database/sql"
	"mp1/pkg/util"
)

var (
	sqlDatabase *sql.DB
	dbLogger    = util.CreateLogger("cmd/db")
)
