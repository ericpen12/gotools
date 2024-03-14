package mysql

import "testing"

func TestGetDB(t *testing.T) {
	db := GetDB("mysql")
	if db != nil {
		t.Log("connect")
	}
}
