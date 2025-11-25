package infra

import (
	"be-lab/common/utils"
)

var Cfg *Config

type Config struct {
	Server *Server
	Mysql  *Mysql
}

type Server struct {
	Host string
	Port string
}

type Mysql struct {
	Host    string
	Port    string
	User    string
	Pwd     string
	DB      string
	Timeout string
}

func NewCfg() {
	Cfg = &Config{
		Server: &Server{
			Host: utils.EnvDefault("SERVER_HOST", "127.0.0.1"),
			Port: utils.EnvDefault("SERVER_PORT", "8000"),
		},
		Mysql: &Mysql{
			Host:    utils.EnvDefault("MYSQL_HOST", "127.0.0.1"),
			Port:    utils.EnvDefault("MYSQL_PORT", "3306"),
			User:    utils.EnvDefault("MYSQL_USER", "lab-manage"),
			Pwd:     utils.EnvDefault("MYSQL_PWD", ""),
			DB:      utils.EnvDefault("MYSQL_DB", "lab-manage"),
			Timeout: utils.EnvDefault("MYSQL_TIMEOUT", "1000"),
		},
	}
}
