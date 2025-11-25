package infra

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	cfg := Cfg.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%sms&charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Pwd, cfg.Host, cfg.Port, cfg.DB, cfg.Timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
