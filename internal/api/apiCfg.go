package api

import (
	"github.com/PhoenixTech2003/go-blog/internal/database"
)

type ApiCfg struct {
	Db *database.Queries
}

var ApiConfig ApiCfg
