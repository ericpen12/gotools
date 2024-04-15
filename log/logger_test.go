package log

import (
	"testing"
)

func TestDebugLog(t *testing.T) {
	Info("test")
	Debug("test")
	//logger.Info(fmt.Sprintf("%v", formatMsg("", 1, 2, "s", "s", errors.New("aa"), []string{"a", "b"})))
}
