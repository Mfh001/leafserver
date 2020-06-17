package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/dbmysql"
	"server/dbredigo"
	"server/models"
	"server/modulegame"
	"server/modulegate"
	"server/modulelogin"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	//start
	dbredigo.Start()
	dbmysql.Start()
	models.CreateTable()

	leaf.Run(
		modulegame.Module,
		modulegate.Module,
		modulelogin.Module,
	)
}
