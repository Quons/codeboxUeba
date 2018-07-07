package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
)

func actUserKeepWeekTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//判断Cursors的值，如果为"init" ，则是初始化任务，否则是日常任务
	if t.Cursors == "init" {
		//初始化任务
		userKeepWeekInitTask(t)
	} else {
		userKeepWeekDailyTask(t)
	}
	wg.Done()
}

func userKeepWeekInitTask(t model.Task) {
	//获取当前时间到14天前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	//遍历前十四天的数据
	for i := -1; i > -14; i-- {
		startTime := currentTime.AddDate(0, 0, i)
		//统计每天数据
		go initTaskStatistic(startTime, currentTime, t)
	}
}

func userKeepWeekInitTaskStatistic(startTime, currentTime time.Time, t model.Task) {
	keepDay := 1
	dayId, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//遍历startTime 到 currentTime之间的
	for currentTime.After(startTime) {
		nextDay := startTime.AddDate(0, 0, 1)
		//查询当天数据
		num, err := postgres.GetUserKeepCount(startTime, nextDay, t)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		//存储数据到mysql
		userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
		mysql.InsertActUserKeepDay(userKeepDay)
		startTime = nextDay
		keepDay++
	}
}

func userKeepWeekDailyTask(t model.Task) {
	//获取当前时间到14天前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	//遍历前十四天的数据
	for i := -1; i > -14; i-- {
		startTime := currentTime.AddDate(0, 0, i)
		//统计每天数据
		go dailyTaskStatistic(startTime, currentTime, t)
	}

}

func userKeepWeekDailyTaskStatistic(startTime, currentTime time.Time, t model.Task) {
	//计算keepday，为starttime到currentTime的差值
	start, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	end, err := strconv.Atoi(currentTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	keepDay := end - start
	//求dayId
	dayId, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//查询gp
	num, err := postgres.GetUserKeepCount(startTime, currentTime, t)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//存储数据到mysql
	userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
	mysql.InsertActUserKeepDay(userKeepDay)
}
