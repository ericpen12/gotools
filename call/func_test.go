package call

import (
	"fmt"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	t.Log(GetFuncName(hello))
}

func hi() {
	fmt.Println("hi")
}
func hello() {
	fmt.Println("hello")
}

func Say(msg string, time int) {
	fmt.Println(msg, time)
}

func TestRegister(t *testing.T) {
	Register(hello, hi)
	for k := range fSet {
		t.Log(k)
	}
}

func TestRegister2(t *testing.T) {
	Register(Say)
	for k := range fSet {
		t.Log(k)
	}
	Call("github.com/ericpen12/gotools/pkg.Say", "ok", 1)
}
