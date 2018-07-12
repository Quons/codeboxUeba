package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestUserKeepWeekInitTask(t *testing.T) {
	rc := make(chan *model.Task)
	go func() {
		select {
		case <-rc:
		}
	}()
	userKeepWeekInitTask(model.Task{Id: 123, ConfigId: 1})
}
func TestUserKeepWeekDailyTask(t *testing.T) {
	userKeepWeekDailyTask(model.Task{Id: 123, ConfigId: 1})
	time.Sleep(1 * time.Second)
}
