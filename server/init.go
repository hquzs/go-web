package server

import (
	"hquzs/go-web/util"
)

func setLevel(level string) {
	util.SetLevel(level)
}

var (
	log = util.NewZeroLog("web", "server")
)
