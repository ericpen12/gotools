package config

import "fmt"

type config struct {
	scene  string
	client Client
}

var cfg *config

func NewConfig(scene string, opts ...Option) {
	cfg = &config{scene: scene}
	for _, opt := range opts {
		opt(cfg)
	}
}

type Client interface {
	Get(key string) (string, error)
	Set(key, value string) error
	BindJson(key string, ptr any) error
	Del(key string) error
}

type Option func(c *config)

func WithClientLocalDB() Option {
	return func(c *config) {
		client, err := newClientDB()
		if err != nil {
			panic("数据库连接失败:" + err.Error())
		}
		c.client = client
	}
}

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
	return c.Get(key)
}

func Set(key, value string) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.Set(key, value)
}

func BindJson(key string, ptr any) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.BindJson(key, ptr)
}

func Del(key string) error {
	c, err := saveClient()
	if err != nil {
		return err
	}
	return c.Del(key)
}
