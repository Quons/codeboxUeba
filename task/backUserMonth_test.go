package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestBackUserMonthStatistics(t *testing.T) {
	backUserMonthStatistics(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, -1, 0), time.Now())
}
