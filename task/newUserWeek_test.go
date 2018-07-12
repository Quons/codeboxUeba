package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestNewUserWeekInsert(t *testing.T) {
	newUserWeekStatistics(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, 0, -7), time.Now())
}
