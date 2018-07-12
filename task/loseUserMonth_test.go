package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestLoseUserMonthStatistics(t *testing.T) {
	loseUserMonthStatistics(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, -1, 0), time.Now())
}
