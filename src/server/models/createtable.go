package models

import "server/dbmysql"

func CreateTable() {
	tablePlayer := dbmysql.Db.HasTable(Player{})
	if !tablePlayer {
		dbmysql.Db.CreateTable(Player{})
	}
}