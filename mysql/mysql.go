package mysql

import (
	"fmt"
	"github.com/ericpen12/gotools/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(name string) *gorm.DB {
	var cfg Config
	err := config.BindJson(name, &cfg)
	if err != nil {
		panic(err)
	}
	db, err := connect(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

type Config struct {
	Username string
	Password string
	Database string
	Host     string
	Port     int
}

func connect(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetCommonDB(name string) *gorm.DB {
	var cfg Config
	err := config.CommonBindJson(name, &cfg)
	if err != nil {
		panic(err)
	}
	db, err := connect(cfg)
	if err != nil {
		panic(err)
	}
	return db
}
