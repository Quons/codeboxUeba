package task

import (
	"testing"
	"time"
	"codeboxUeba/model"
	"codeboxUeba/mysql"
)

func TestDailyTaskStatistic(t *testing.T) {
	mysql.Init()

	dailyTaskStatistic(time.Now(),time.Now().AddDate(0,0,1),model.Task{})

}