package mysql

import (
	"testing"
	"codeboxUeba/model"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

func init() {
	Init()
}

func TestInsertActUserDay(t *testing.T) {
	actUserDay := &model.ActUserDay{Num: 12, DayId: 20060102, ConfigId: 11111}
	err := InsertActUserDay(actUserDay)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertActUserWeek(t *testing.T) {
	actUserWeek := &model.ActUserWeek{Num: 123, ConfigId: 0, StartDay: 20060102, EndDay: 20060103, WeekId: 200601}
	err := InsertActUserWeek(actUserWeek)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertActUserMonth(t *testing.T) {
	actUserMonth := &model.ActUserMonth{Num: 123, ConfigId: 0, MonthId: 200601}
	err := InsertActUserMonth(actUserMonth)
	if err != nil {
		t.Error(err)
	}
}

func TestInit(t *testing.T) {
	bs := sha256.Sum256([]byte("abcd"))
	fmt.Println(bs)

	hash := sha256.New()
	hash.Write([]byte("abcd"))
	md := hash.Sum(nil)
	fmt.Println(string(md[:]))
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
}
