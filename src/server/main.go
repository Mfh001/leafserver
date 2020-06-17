package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game"
	"server/gate"
	"server/login"
	"server/redigo"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	if err := redigo.Setup(); err == nil {
		log.Release("create redis pool")
	}

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
