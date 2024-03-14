package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig(t *testing.T) {
	t.Log(viper.Get("mysql.username"))

}
