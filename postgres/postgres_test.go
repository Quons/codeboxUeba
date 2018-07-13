package postgres

import (
	"testing"
	"time"
	"codeboxUeba/model"
	"codeboxUeba/mysql"
)

func init() {
	mysql.Init()
}

func TestQueryCount(t *testing.T) {
	sql := "SELECT count(1) FROM dw_requestlog where platfrom=$1"
	num, err := QueryCount(sql, "ios")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}

func TestGetUserKeepCount(t *testing.T) {
	now := time.Now()
	num, err := GetUserKeepCount(now.AddDate(0, 0, -7), now.AddDate(0, 0, -6), now.AddDate(0, 0, -3),
		now.AddDate(0, 0, -4), model.Task{ConfigId: 1})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}

func TestGetGpCount(t *testing.T) {
	num, err := GetGpCount(1, time.Now().AddDate(0, 0, -2), time.Now().AddDate(0, 0, -1))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}

func TestFunnelCount(t *testing.T) {
	num, err := FunnelCount([]string{"/user/login", "/user/life"}, time.Now().AddDate(0, 0, -2), time.Now().AddDate(0, 0, -1))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}
