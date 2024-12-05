package config

import "fmt"

type config struct {
	scene  string
	client Client
}

var cfg *config

func NewConfig(scene string, opts ...Option) {
	client, err := newClientDB()
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	cfg = &config{scene: scene, client: client}
	for _, opt := range opts {
		opt(cfg)
	}
}

type Client interface {
	Get(key, scene string) (string, error)
	Set(key, value, scene string) error
	BindJson(key string, ptr any, scene string) error
	Del(key, scene string) error
}

type Option func(c *config)

func saveClient() (Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cfg 没有初始化")
	}
	if cfg.client == nil {
		return nil, fmt.Errorf("client 没有初始化")
	}
	return cfg.client, nil
}

func Get(key string) (string, error) {
	c, err := saveClient()
	if err != nil {
		return "", err
	}
	return c.Get(key, cfg.scene)
}

func Set(key, value string) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.Set(key, value, cfg.scene)
}

func BindJson(key string, ptr any) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.BindJson(key, ptr, cfg.scene)
}

const commonScene = "common-config"

func CommonBindJson(key string, ptr any) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.BindJson(key, ptr, commonScene)
}

func Del(key string) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.Del(key, cfg.scene)
}

func AppName() string {
	return cfg.scene
}
