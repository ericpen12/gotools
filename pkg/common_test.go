package pkg

import (
	"testing"
	"time"
)

func TestPrintRunCost(t *testing.T) {
	defer PrintRunCost()
	time.Sleep(time.Second)
}
