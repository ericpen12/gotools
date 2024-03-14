package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initResource() {
	list := []func() error{initMysql}
	for _, fn := range list {
		err := fn()
		if err != nil {
			panic(err)
		}
	}
}

var DB *gorm.DB

type mysqlConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     int
}

func initMysql() error {
	if viper.Get("mysql") == nil {
		return nil
	}
	var config mysqlConfig
	err := viper.UnmarshalKey("mysql", &config)
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("数据库：%s 已连接\n", viper.Get("mysql.host"))
	return nil
}
