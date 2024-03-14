package pkg

import (
	"testing"
	"time"
)

func TestPrintRunCost(t *testing.T) {
	defer PrintRunCost()
	timeNow := time.Now().Add(-60 * time.Second)

	time.Since(timeNow).Seconds()
}

func TestGetCurrentAppName(t *testing.T) {
	app := GetCurrentAppName()
	t.Log(app)
}
