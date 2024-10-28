package cookiefile

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
	"time"
)

func New(filename string, expireTime time.Duration) *Cookie {
	return &Cookie{
		filename:   filename,
		expireTime: expireTime,
	}
}

type Cookie struct {
	filename   string
	expireTime time.Duration
}

func (c *Cookie) Get() string {
	fileInfo, err := os.Stat(c.filename)
	if err != nil {
		return ""
	}
	if fileInfo.ModTime().Before(time.Now().Add(-c.expireTime)) {
		_ = os.Remove(c.filename)
		return ""
	}
	b, _ := os.ReadFile(c.filename)
	return string(b)
}

func (c *Cookie) Set(value string) {
	err := os.WriteFile(c.filename, []byte(value), os.ModePerm)
	fmt.Println(err)

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
