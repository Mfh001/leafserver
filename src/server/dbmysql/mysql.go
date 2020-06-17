package dbmysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/name5566/leaf/log"
	"server/conf"
)

var Db *gorm.DB

func Start() {
	var err error
	Db, err = gorm.Open(conf.Server.MysqlType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Server.MysqlUser,
		conf.Server.MysqlPassword,
		conf.Server.MysqlHost,
		conf.Server.MysqlName))

	if err != nil {
		log.Fatal("mysql Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.Server.MysqlTablePrefix + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(conf.Server.MysqlMaxIdle)
	Db.DB().SetMaxOpenConns(conf.Server.MysqlMaxOpen)
	log.Release("mysql setup finished")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer Db.Close()
}