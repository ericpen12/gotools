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

func (c *Cookie) Get() (string, error) {
	cookie := c.getCache()
	if cookie != "" {
		return cookie, nil
	}
	c1 := time.After(time.Second * 10)
	for {
		select {
		case <-c1:
			return "", fmt.Errorf("未获取到cookie信息")
		default:
			time.Sleep(time.Second)
			cookie = getFromClipboard()
			if cookie != "" {
				c.set(cookie)
				return cookie, nil
			}
		}
	}
}

func (c *Cookie) getCache() string {
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

func (c *Cookie) Expire() {
	_ = os.Remove(c.filename)
}

func (c *Cookie) set(value string) {
	_ = os.WriteFile(c.filename, []byte(value), os.ModePerm)
}

func getFromClipboard() string {
	text, err := clipboard.ReadAll()
	if err != nil {
		return ""
	}
	if !strings.Contains(text, "=") {
		return ""
	}
	return text
}
