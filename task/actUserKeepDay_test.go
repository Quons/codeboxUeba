package task

import (
	"testing"
	"time"
	"codeboxUeba/model"
	"codeboxUeba/mysql"
)

func init() {
	mysql.Init()
}

func TestDailyTaskStatistic(t *testing.T) {
	dailyTaskStatistic(time.Now().AddDate(0, 0, -3), time.Now().AddDate(0, 0, -2), model.Task{Id: 1, ConfigId: 1})
}

func TestUserKeepInitTask(t *testing.T) {
	rc := make(chan *model.Task)
	go func() {
		select {
		case <-rc:
		}
	}()

	userKeepInitTask(model.Task{Id: 1, ConfigId: 1})
}
func TestUserKeepDailyTask(t *testing.T) {
	userKeepDailyTask(model.Task{Id: 1, ConfigId: 1})
	time.Sleep(1 * time.Second)
}
