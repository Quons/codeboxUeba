package mysql

import (
	"testing"
	"codeboxUeba/model"
)

func TestInsertBackUserWeek(t *testing.T) {
	backUserWeek := &model.BackUserWeek{ConfigId: 1, Num: 123, StartDay: 20060102, EndDay: 20060103, WeekId: 201827}
	err := InsertBackUserWeek(backUserWeek)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInsertBackUserMonth(t *testing.T) {
	backUserMonth := &model.BackUserMonth{Num: 123, ConfigId: 1, MonthId: 201807}
	err := InsertBackUserMonth(backUserMonth)
	if err != nil {
		t.Error(err)
	}
}
