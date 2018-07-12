package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestNewUserMonthStatistics(t *testing.T) {
	newUserMonthStatistics(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, -1, 0), time.Now())
}
