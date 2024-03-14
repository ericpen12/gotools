package config

import (
	"fmt"
	"github.com/ericpen12/gotools/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	list := []func() error{
		initViper,
		initMysql,
	}
	for _, fn := range list {
		err := fn()
		if err != nil {
			panic(err)
		}
	}
}

func initViper() error {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
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
	log.Infof("数据库：%s 已连接\n", viper.Get("mysql.host"))
	return nil
}
