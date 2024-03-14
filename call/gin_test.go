package call

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func NoArgs() string {
	return "hi"
}

func Json(m map[string]any) any {
	fmt.Println(m)
	return m
}

type in struct {
	Name string
	Age  int `json:"a"`
}

func Struct(m in) in {
	fmt.Println(m)
	return m
}

func TestServer(t *testing.T) {
	r := gin.Default()
	// 注册服务
	//RegisterHandler(r, "quickstart", &service{})
	Register(NoArgs, Json, Struct)
	r.POST("quickstart", Driver)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
