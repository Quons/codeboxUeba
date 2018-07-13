package task

import (
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
	"math"
	"fmt"
)

func actUserKeepDayTask(t model.Task) {
	if t.FromDate != "" {
		//初始化任务
		userKeepInitTask(t)
	} else {
		userKeepDailyTask(t)
	}

}

func userKeepInitTask(t model.Task) {
	//获取当前时间到14天前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, 0, -1)

	//遍历前十四天的数据
	for i := 0; i < 14; i++ {
		startTime := currentTime.AddDate(0, 0, -i)
		//统计每天数据
		go initTaskStatistic(startTime, currentTime, t)
	}
}

func initTaskStatistic(startTime, currentTime time.Time, t model.Task) {
	keepDay := 0
	dayId, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	//遍历startTime 到 currentTime之间的
	tmpTime := startTime
	for currentTime.After(tmpTime) || currentTime == tmpTime {
		nextDay := tmpTime.AddDate(0, 0, 1)
		//查询当天数据
		num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 1), nextDay, nextDay.AddDate(0, 0, 1), t)
		if err != nil {
			RecordFailTask(startTime, nextDay, &t)
			log.LogError(err.Error())
			return
		}
		//存储数据到mysql
		userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
		err = mysql.InsertActUserKeepDay(userKeepDay)
		if err != nil {
			RecordFailTask(startTime, nextDay, &t)
			log.LogError(err.Error())
			return
		}
		fmt.Printf("initTaskSatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
		tmpTime = nextDay
		keepDay++
	}
}

func userKeepDailyTask(t model.Task) {
	//获取当前时间到14天前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, 0, -1)

	//遍历前十四天的数据
	for i := 0; i < 14; i++ {
		startTime := currentTime.AddDate(0, 0, -i)
		//统计每天数据
		go dailyTaskStatistic(startTime, currentTime, t)
	}
}

func dailyTaskStatistic(startTime, currentTime time.Time, t model.Task) {
	duration := currentTime.Sub(startTime)
	keepDay := int(math.Floor((duration.Hours() / 24) + 0.5))
	dayId, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	//查询gp
	num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 1), currentTime, currentTime.AddDate(0, 0, 1), t)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	//存储数据到mysql
	userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
	err = mysql.InsertActUserKeepDay(userKeepDay)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(startTime, currentTime, &t)
		return
	}
	fmt.Printf("dailyTaskSatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
}
