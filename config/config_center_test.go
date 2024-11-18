package config

import "testing"

func TestConfigCenter(t *testing.T) {
	NewConfig(WithClientLocalDB())
	ret, err := Get("1")
	if err != nil {
		t.Log(err)
	}
	t.Log(ret)

	err = Set("1", "{\"a\":1}")
	if err != nil {
		t.Log(err)
	}
	ret, _ = Get("1")
	_ = Del("1")
	t.Log(ret)
}
