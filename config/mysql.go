package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type configModel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Scene       string `json:"scene"`
	Config      string `json:"config"`
}

func (c *configModel) TableName() string {
	return "config"
}

type clientDB struct {
	db *gorm.DB
}

func (c *clientDB) Get(key string) (string, error) {
	var result configModel
	err := c.db.Take(&result, "name=? and scene=?", key, cfg.scene).Error
	return result.Config, err
}

func (c *clientDB) Set(key, value string) error {
	ret, _ := c.Get(key)
	if len(ret) == 0 {
		return c.db.Create(&configModel{
			Name:   key,
			Scene:  cfg.scene,
			Config: value,
		}).Error
	}
	return c.db.Model(&configModel{}).
		Where("name=? and scene=?", key, cfg.scene).
		Update("config", value).Error
}

func (c *clientDB) BindJson(key string, ptr any) error {
	return c.db.Take(ptr, "name=? and scene=?", key, cfg.scene).Error
}

func (c *clientDB) Del(key string) error {
	return c.db.Model(&configModel{}).Delete(nil, "name=? and scene=?", key, cfg.scene).Error
}

func newClientDB() (*clientDB, error) {
	dsn := "configcenter:configcenter12@tcp(localhost:3306)/project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return &clientDB{db: db}, nil
}
