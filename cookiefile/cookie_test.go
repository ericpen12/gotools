package cookiefile

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New("./aa.cookie", time.Second*10)
	//s.Set("ddd")
	ret := s.Get()
	t.Log(ret)
}
