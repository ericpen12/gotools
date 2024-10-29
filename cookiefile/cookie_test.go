package cookiefile

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	s := New("./aa.cookie", time.Second*10)
	//s.Set("ddd")
	ret, err := s.Get()
	t.Log(err)
	t.Log(ret)
}
