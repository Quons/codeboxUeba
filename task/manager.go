package task

import (
	"codeboxUeba/model"
	. "codeboxUeba/conf"
	"sync"
	"time"
	"codeboxUeba/log"
	"codeboxUeba/utils"
	"fmt"
)

func TasksFactory(taskName string) (f func(wg *sync.WaitGroup, rc chan *model.Task, task model.Task)) {
	switch taskName {
	case ActUserWeek:
		return actUserWeekTask
	case ActUserDay:
		return actUserDayTask
	case ActUserMonth:
		return actUserMonthTask
	case NewUserDay:
		return newUserDayTask
	case NewUserWeek:
		return newUserWeekTask
	case NewUserMonth:
		return newUserMonthTask
	case ActUserKeepDay:
		return actUserKeepDayTask
	case ActUserKeepWeek:
		return actUserKeepWeekTask
	case ActUserKeepMonth:
		return actUserKeepMonthTask
	case BackUserWeek:
		return backUserWeekTask
	default:
		return nil
	}
}

func dayStatistic(wg *sync.WaitGroup, rc chan *model.Task, t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time)) {
	//获取cursor到当前时间的时间列表 todo 添加指定时间段的功能
	//fmt.Println("cursor.......",t.Cursors)
	fromDate, err := time.Parse("20060102", t.Cursors)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	nowDate := time.Now()
	for {
		toDate := fromDate.AddDate(0, 0, 1)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//更新执行状态 数据库或者文件，暂时使用文件进行记录
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}
		fmt.Println("fromdate:",fromDate," todate:",toDate,"nowDate:",nowDate)
		go f(t, fromDate, toDate)
		fromDate = toDate
	}
}

func monthStatistic(wg *sync.WaitGroup, rc chan *model.Task, t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time)) {
	//获取cursor到当前时间的时间列表
	fromDate, err := time.Parse("20060102", t.Cursors)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	// 获取本月的一号 todo 添加指定时间段的功能
	fromDate, err = time.Parse("200601", fromDate.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	nowDate := time.Now()
	for {
		toDate := fromDate.AddDate(0, 1, 0)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}

		go f(t, fromDate, toDate)
		//actUserDayChan <- &model.ActUserDayStatistic{t, fromDate, toDate}
		fromDate = toDate
	}
}

func weekStatistic(wg *sync.WaitGroup, rc chan *model.Task, t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time)) {
	//获取cursor到当前时间的时间列表 todo 添加指定时间段的功能
	fromDate, err := time.Parse("20060102", t.Cursors)
	utils.CheckError(err)

	//获取本周的周一
	weekDay := fromDate.Weekday()
	for weekDay != time.Monday {
		fromDate = fromDate.AddDate(0, 0, -1)
		weekDay = fromDate.Weekday()
	}

	nowDate := time.Now()
	for {
		toDate := fromDate.AddDate(0, 0, 7)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}

		go f(t, fromDate, toDate)
		//actUserDayChan <- &model.ActUserDayStatistic{t, fromDate, toDate}
		fromDate = toDate
	}
}
