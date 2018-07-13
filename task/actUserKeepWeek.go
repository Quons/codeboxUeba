package task

import (
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
	"fmt"
)

func actUserKeepWeekTask(t model.Task) {
	if t.FromDate != "" {
		//初始化任务
		userKeepWeekInitTask(t)
	} else {
		userKeepWeekDailyTask(t)
	}
}

func userKeepWeekInitTask(t model.Task) {
	//获取当前时间到7周前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, 0, -7)

	//获取周一
	for currentTime.Weekday() != time.Monday {
		currentTime = currentTime.AddDate(0, 0, -1)
	}
	//遍历前7周的数据
	for i := 0; i < 7; i++ {
		startTime := currentTime.AddDate(0, 0, -7*i)
		//统计每天数据
		go userKeepWeekInitTaskStatistic(startTime, currentTime, t)

	}
}

func userKeepWeekInitTaskStatistic(startTime, currentTime time.Time, t model.Task, ) {
	keepWeek := 0
	//weekId为当年的第几周
	year, week := startTime.ISOWeek()
	WeekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}

	//遍历startTime 到 currentTime之间的
	tmpTime := startTime
	for currentTime.After(tmpTime) || currentTime == tmpTime {
		nextWeek := tmpTime.AddDate(0, 0, 7)
		//查询当天数据
		num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 7), nextWeek, nextWeek.AddDate(0, 0, 7), t)
		if err != nil {
			log.LogError(err.Error())
			RecordFailTask(startTime, nextWeek, &t)
			return
		}
		//存储留存数据到mysql
		userKeepWeek := &model.ActUserKeepWeek{WeekId: WeekId, KeepWeek: keepWeek, Num: num, ConfigId: t.ConfigId}
		err = mysql.InsertActUserKeepWeek(userKeepWeek)
		if err != nil {
			log.LogError(err.Error())
			RecordFailTask(startTime, nextWeek, &t)
			return
		}
		fmt.Printf("userKeepWeekInitTaskStatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
		tmpTime = nextWeek
		keepWeek++
	}
}

func userKeepWeekDailyTask(t model.Task) {
	//获取当前时间到7周前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, 0, -7)

	//获取周一
	for currentTime.Weekday() != time.Monday {
		currentTime = currentTime.AddDate(0, 0, -1)
	}

	//遍历前7周的数据
	for i := 0; i < 7; i++ {
		startTime := currentTime.AddDate(0, 0, -7*i)
		//统计每天数据
		go userKeepWeekDailyTaskStatistic(startTime, currentTime, t)

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
		RecordFailTask(startTime, currentTime, &t)
		return
	}

	num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 7), currentTime, currentTime.AddDate(0, 0, 7), t)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	userKeepWeek := &model.ActUserKeepWeek{WeekId: WeekId, KeepWeek: keepWeek, Num: num, ConfigId: t.ConfigId}
	err = mysql.InsertActUserKeepWeek(userKeepWeek)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	fmt.Printf("userKeepWeekDailyTaskStatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
}
