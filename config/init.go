package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	ErrConfigNotFound  = fmt.Errorf("配置不存在")
	ErrConfigStructure = fmt.Errorf("配置结构错误")
)

func init() {
	err := initViper()
	if err != nil {
		panic(err)
	}
}

func initViper() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // optionally look for config in the working directory // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".") // optionally look for config in the working directory
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(u + "/Develop/.setting") // call multiple times to add many search paths
	err = viper.ReadInConfig()                   // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
}

type MysqlConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     int
}

func Mysql(configName string) (*MysqlConfig, error) {
	if viper.Get(configName) == nil {
		return nil, ErrConfigNotFound
	}
	var config MysqlConfig
	err := viper.UnmarshalKey(configName, &config)
	if err != nil {
		return nil, ErrConfigStructure
	}
	return &config, nil
}

type LogConfig struct {
	Level string
}

func Log() LogConfig {
	var c LogConfig

	return c
}
