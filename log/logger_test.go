package log

import (
	"github.com/ericpen12/gotools/config"
	"testing"
)

func TestDebugLog(t *testing.T) {
	config.NewConfig("gotools", config.WithClientLocalDB())
	Init()
	Info("test")
	Debug("test")
	//logger.Info(fmt.Sprintf("%v", formatMsg("", 1, 2, "s", "s", errors.New("aa"), []string{"a", "b"})))
}
