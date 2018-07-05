package task

import (
	"uebaDataJob/model"
	. "uebaDataJob/conf"
	"sync"
)

func TasksFactory(taskName string) (f []func(wg *sync.WaitGroup, rc chan *model.Task, task model.Task)) {
	switch taskName {
	case ActUser:
		return []func(wg *sync.WaitGroup, rc chan *model.Task, task model.Task){
			actUserDayTask,
			actUserWeekTask,
			actUserMonthTask,
		}
	case NewUser:
		return []func(wg *sync.WaitGroup, rc chan *model.Task, task model.Task){
			newUserDayTask,
			newUserWeekTask,
			newUserMonthTask,
		}
	default:
		return nil
	}
}
