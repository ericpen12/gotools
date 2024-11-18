package cookiefile

import (
	"fmt"
	"github.com/atotto/clipboard"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func New(link string, expireTime time.Duration) *Cookie {
	u, _ := url.Parse(link)
	return &Cookie{
		filename:   u.Host + ".cookie",
		url:        link,
		expireTime: expireTime,
	}
}

type Cookie struct {
	filename   string
	url        string
	expireTime time.Duration
}

const ReadyReadFromClipboard = "ReadyRead"

func (c *Cookie) Get() (string, error) {
	cookie := c.getCache()
	if cookie != "" {
		return cookie, nil
	}
	_ = exec.Command(`open`, c.url).Start()
	_ = clipboard.WriteAll(ReadyReadFromClipboard)
	c1 := time.After(time.Minute)
	for {
		select {
		case <-c1:
			return "", fmt.Errorf("未获取到cookie信息")
		default:
			time.Sleep(time.Second)
			cookie, err := getFromClipboard()
			if err != nil {
				return "", err
			}
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

func getFromClipboard() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	if text == ReadyReadFromClipboard {
		return "", nil
	}
	if !strings.Contains(text, "=") {
		return "", fmt.Errorf("cookie格式有误")
	}
	return text, nil
}
