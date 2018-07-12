package mysql

import (
	"testing"
	"codeboxUeba/model"
)

func TestInsertNewUserDay(t *testing.T) {
	newUserDay := &model.NewUserDay{Num: 123, ConfigId: 1, DayId: 20180711}
	err := InsertNewUserDay(newUserDay)
	if err != nil {
		t.Log(err)
		return
	}
}

func TestInsertNewUserWeek(t *testing.T) {
	newUserWeek := &model.NewUserWeek{Num: 123, ConfigId: 1, StartDay: 20180710, EndDay: 20180711}
	err := InsertNewUserWeek(newUserWeek)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInsertNewUserMonth(t *testing.T) {
	newUserMonth := &model.NewUserMonth{Num: 123, MonthId: 201807, ConfigId: 1}
	err := InsertNewUserMonth(newUserMonth)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQueryNewUserCurrentWeek(t *testing.T) {
	num, err := QueryNewUserCurrentWeek(1, 20180727)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}
func TestQueryNewUserCurrentMonth(t *testing.T) {
	num, err := QueryNewUserCurrentMonth(1, 201807)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(num)
}
