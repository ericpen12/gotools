package config

import (
	"encoding/json"
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

func (c *clientDB) Get(key, scene string) (string, error) {
	var result configModel
	err := c.db.Take(&result, "name=? and scene=?", key, scene).Error
	return result.Config, err
}

func (c *clientDB) Set(key, value, scene string) error {
	ret, _ := c.Get(key, scene)
	if len(ret) == 0 {
		return c.db.Create(&configModel{
			Name:   key,
			Scene:  scene,
			Config: value,
		}).Error
	}
	return c.db.Model(&configModel{}).
		Where("name=? and scene=?", key, scene).
		Update("config", value).Error
}

func (c *clientDB) BindJson(key string, ptr any, scene string) error {
	var res configModel
	err := c.db.Take(&res, "name=? and scene=?", key, scene).Error
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(res.Config), ptr)
}

func (c *clientDB) Del(key, scene string) error {
	return c.db.Model(&configModel{}).Delete(nil, "name=? and scene=?", key, scene).Error
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
