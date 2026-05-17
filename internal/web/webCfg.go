package web

import (
	"github.com/PhoenixTech2003/go-blog/internal/database"
)

type WebCfg struct {
	Db *database.Queries
}

var WebConfig WebCfg
