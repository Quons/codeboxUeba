package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestUserKeepMonthInitTask(t *testing.T) {
	rc := make(chan *model.Task)
	go func() {
		select {
		case <-rc:
		}
	}()
	userKeepMonthInitTask(model.Task{Id: 1, ConfigId: 1})
	time.Sleep(1 * time.Second)
}

func TestUserKeepMonthDailyTask(t *testing.T) {
	rc := make(chan *model.Task)
	go func() {
		select {
		case <-rc:
		}
	}()
	userKeepMonthDailyTask(model.Task{Id: 1, ConfigId: 1})
	time.Sleep(1 * time.Second)
}
