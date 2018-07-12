package mysql

import (
	"testing"
	"codeboxUeba/model"
)

func TestInsertLoseUserWeek(t *testing.T) {

	loseUserWeek := &model.LoseUserWeek{Num: 123, WeekId: 201827, ConfigId: 1}
	err := InsertLoseUserWeek(loseUserWeek)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInsertLoseUserMonth(t *testing.T) {
	loseUserMonth := &model.LoseUserMonth{Num: 123, MonthId: 201807, ConfigId: 1}
	err := InsertLoseUserMonth(loseUserMonth)
	if err != nil {
		t.Error(err)
		return
	}
}
