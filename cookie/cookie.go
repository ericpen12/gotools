package cookie

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func NewFile(filename string) *Cookie {
	return &Cookie{
		filename: filename,
	}
}

type Cookie struct {
	filename string
	buf      map[string]string
}

func (c *Cookie) Get(key string) string {
	b, _ := os.ReadFile(c.filename)
	_ = json.Unmarshal(b, &c.buf)
	return c.buf[key]
}

func (c *Cookie) Set(key, value string) {
	c.buf[key] = value
	b, _ := json.Marshal(c.buf)
	_ = os.WriteFile(c.filename, b, os.ModePerm)
}

func GetFromClipboard() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	if !strings.Contains(text, "=") {
		return "", fmt.Errorf("未获取到cookie信息")
	}
	return text, nil
}
