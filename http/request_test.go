package http

import "testing"

func TestCookie(t *testing.T) {
	res, err := EasyRequest(WithUrl("https://baidu.com")).Do()
	if err != nil {
		t.Log(err)
	}
	t.Log(res.Cookie())
}
