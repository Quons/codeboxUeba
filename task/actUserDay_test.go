package task

import (
	"testing"
	"codeboxUeba/mysql"
	"codeboxUeba/model"
	"time"
)

func init() {
	mysql.Init()
}

func TestActUserInsert(t *testing.T) {
	ActUserDayInsert(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, 0, -3), time.Now().AddDate(0, 0, -2))
}
