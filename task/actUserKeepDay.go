package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
	"math"
	"fmt"
)

func actUserKeepDayTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//判断Cursors的值，不为当前时间，则是初始化任务，否则是日常任务
	if t.Cursors != time.Now().Format("20060102") {
		//初始化任务
		userKeepInitTask(t, rc)
	} else {
		userKeepDailyTask(t)
	}
	wg.Done()
}

func userKeepInitTask(t model.Task, rc chan *model.Task) {
	//获取当前时间到14天前的时间列表
	currentTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, 0, -1)
	wg := &sync.WaitGroup{}

	//遍历前十四天的数据
	for i := 0; i < 14; i++ {
		wg.Add(1)
		startTime := currentTime.AddDate(0, 0, -i)
		//统计每天数据
		go initTaskStatistic(startTime, currentTime, t, wg)
	}
	wg.Wait()
	t.Cursors = currentTime.Format("20060102")
	rc <- &t
}

func initTaskStatistic(startTime, currentTime time.Time, t model.Task, wg *sync.WaitGroup) {
	keepDay := 0
	dayId, err := strconv.Atoi(startTime.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//遍历startTime 到 currentTime之间的
	tmpTime := startTime
	for currentTime.After(tmpTime) || currentTime == tmpTime {
		nextDay := tmpTime.AddDate(0, 0, 1)
		//查询当天数据
		num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 1), nextDay, nextDay.AddDate(0, 0, 1), t)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		//存储数据到mysql
		userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
		err = mysql.InsertActUserKeepDay(userKeepDay)
		if err != nil {
			log.LogError(err.Error())
		}
		fmt.Printf("initTaskSatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
		tmpTime = nextDay
		keepDay++
	}
	wg.Done()

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
		return
	}
	//查询gp
	num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 0, 1), currentTime, currentTime.AddDate(0, 0, 1), t)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//存储数据到mysql
	userKeepDay := &model.ActUserKeepDay{DayId: dayId, KeepDay: keepDay, Num: num, ConfigId: t.ConfigId}
	err = mysql.InsertActUserKeepDay(userKeepDay)
	if err != nil {
		log.LogError(err.Error())
	}
	fmt.Printf("dailyTaskSatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)

}
