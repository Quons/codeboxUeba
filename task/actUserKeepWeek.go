package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
	"fmt"
)

func actUserKeepWeekTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//判断Cursors的值，如果为"init" ，则是初始化任务，否则是日常任务
	if t.Cursors != time.Now().Format("20060102") {
		//初始化任务
		userKeepWeekInitTask(t, rc)
	} else {
		userKeepWeekDailyTask(t)
	}
	wg.Done()
}

func userKeepWeekInitTask(t model.Task, rc chan *model.Task) {
	//获取当前时间到7周前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	//获取周一
	for currentTime.Weekday() != time.Monday {
		fmt.Println(currentTime)
		currentTime = currentTime.AddDate(0, 0, -1)
	}
	wg := &sync.WaitGroup{}
	//遍历前7周的数据
	for i := 1; i < 7; i++ {
		wg.Add(1)
		startTime := currentTime.AddDate(0, 0, -7*i)
		//统计每天数据
		go userKeepWeekInitTaskStatistic(startTime, currentTime, t, wg)
	}
	wg.Wait()
	//todo 暂时使用当前时间存档
	t.Cursors = time.Now().Format("20060102")
	rc <- &t
}

func userKeepWeekInitTaskStatistic(startTime, currentTime time.Time, t model.Task, wg *sync.WaitGroup) {
	keepWeek := 0
	//weekId为当年的第几周
	year, week := startTime.ISOWeek()
	WeekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	//遍历startTime 到 currentTime之间的
	tmpTime := startTime
	for currentTime.After(tmpTime) {
		nextWeek := tmpTime.AddDate(0, 0, 7)
		//查询当天数据
		num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 7), nextWeek, nextWeek.AddDate(0, 0, 7), t)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		//存储留存数据到mysql
		userKeepWeek := &model.ActUserKeepWeek{WeekId: WeekId, KeepWeek: keepWeek, Num: num, ConfigId: t.ConfigId}
		mysql.InsertActUserKeepWeek(userKeepWeek)
		tmpTime = nextWeek
		keepWeek++
	}
	wg.Done()
}

func userKeepWeekDailyTask(t model.Task) {
	//获取当前时间到7周前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	//获取周一
	for currentTime.Weekday() != time.Monday {
		currentTime = currentTime.AddDate(0, 0, -1)
	}

	//遍历前7周的数据
	for i := 1; i < 7; i++ {
		startTime := currentTime.AddDate(0, 0, -7)
		//统计每天数据
		go userKeepWeekDailyTaskStatistic(startTime, currentTime, t)
		currentTime = startTime
	}

}

func userKeepWeekDailyTaskStatistic(startTime, currentTime time.Time, t model.Task) {
	_, cw := currentTime.ISOWeek()
	//weekId为当年的第几周
	sy, sw := startTime.ISOWeek()
	//计算keepWeek，为当前时间到开始时间的差值
	keepWeek := cw - sw

	WeekId, err := strconv.Atoi(fmt.Sprintf("%v%v", sy, sw))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 7), currentTime, currentTime.AddDate(0, 0, 7), t)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	userKeepWeek := &model.ActUserKeepWeek{WeekId: WeekId, KeepWeek: keepWeek, Num: num, ConfigId: t.ConfigId}
	mysql.InsertActUserKeepWeek(userKeepWeek)
}
