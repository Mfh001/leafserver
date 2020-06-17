package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"time"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string

	//redis
	RedisHost string
	RedisPassword string
	RedisMaxIdle int
	RedisMaxActive int
	RedisIdleTimeout time.Duration

	//mysql
	MysqlType string
	MysqlHost string
	MysqlUser string
	MysqlPassword string
	MysqlName string
	MysqlTablePrefix string
	MysqlMaxIdle int
	MysqlMaxOpen int
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if data == nil {
		data, err = ioutil.ReadFile("../conf/server.json")
	}
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
