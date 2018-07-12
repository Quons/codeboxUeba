package mysql

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestInsertActUserKeepDay(t *testing.T) {
	actUserKeepDay := &model.ActUserKeepDay{Num: 123, ConfigId: 0, DayId: 123}
	err := InsertActUserKeepDay(actUserKeepDay)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertActUserKeepWeek(t *testing.T) {
	//todo startday和endDay 不需要？
	actUserKeepWeek := &model.ActUserKeepWeek{Num: 123, ConfigId: 0, WeekId: 200602}
	err := InsertActUserKeepWeek(actUserKeepWeek)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertActUserKeepMonth(t *testing.T) {
	actUserKeepMonth := &model.ActUserKeepMonth{Num: 123, ConfigId: 1, MonthId: 200601}
	err := InsertActUserKeepMonth(actUserKeepMonth)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryUserKeepPreWeek(t *testing.T) {
	var confId int64 = 1
	fromDate := time.Now()
	num, err := QueryUserKeepPreWeek(confId, fromDate)
	if err != nil {
	}
	t.Log("preWeek user keep week num:", num)
}

func TestQueryUserKeepPreMonth(t *testing.T) {
	var confId int64 = 1
	num, err := QueryUserKeepPreMonth(confId, time.Now())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("preMonth user keep num:", num)

}
