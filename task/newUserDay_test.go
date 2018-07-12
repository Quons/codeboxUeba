package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestNewUserDayInsert(t *testing.T) {
	newUserDayInsert(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, 0, -2), time.Now())
}
