package mysql

import (
	"fmt"
	"github.com/ericpen12/gotools/config"
	"github.com/ericpen12/gotools/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(name string) *gorm.DB {
	db, err := connect(name)
	if err != nil {
		panic(err)
	}
	return db
}

func connect(name string) (*gorm.DB, error) {
	cfg, err := config.Mysql(name)
	if err != nil {
		return nil, err
	}
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
	log.Debugf("数据库：%s 已连接", cfg.Host)
	return db, nil
}
