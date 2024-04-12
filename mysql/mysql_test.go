package mysql

import "testing"

func Test_initMysql(t *testing.T) {
	if DB != nil {
		t.Log("connect")
	}
}
