package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestActUserMonthStatistics(t *testing.T) {
	actUserMonthStatistics(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, -0, 0))
}
