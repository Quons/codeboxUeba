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

func actUserKeepMonthTask(t model.Task) {
	//判断Cursors的值，如果为当前时间，则是日常任务，否则是初始化任务
	if t.FromDate != "" {
		//初始化任务
		userKeepMonthInitTask(t)
	} else {
		userKeepMonthDailyTask(t)
	}
}

func userKeepMonthInitTask(t model.Task) {
	//获取当当月一号
	currentTime, err := time.Parse("200601", time.Now().Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	currentTime = currentTime.AddDate(0, -1, 0)

	//遍历前6个月的数据
	for i := 0; i < 7; i++ {
		startTime := currentTime.AddDate(0, -i, 0)
		//统计每天数据
		go func() {
			result := userKeepMonthInitTaskStatistic(startTime, currentTime, t)
			if result == ErrorCode {
				mysql.FailRecord(startTime.Format("200601"), t.Id)

			}
		}()
	}

}

func userKeepMonthInitTaskStatistic(startTime, currentTime time.Time, t model.Task) int {
	keepMonth := 0
	//遍历startTime 到 currentTime之间的
	tmpTime := startTime
	for currentTime.After(tmpTime) || currentTime == tmpTime {
		nextMonth := tmpTime.AddDate(0, 1, 0)
		num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 1, 0), nextMonth, nextMonth.AddDate(0, 1, 0), t)
		if err != nil {
			log.LogError(err.Error())
			return ErrorCode
		}
		//存储数据到mysql
		monthId, err := strconv.Atoi(startTime.Format("200601"))
		if err != nil {
			log.LogError(err.Error())
			return ErrorCode
		}
		userKeepMonth := &model.ActUserKeepMonth{MonthId: monthId, KeepMonth: keepMonth, Num: num, ConfigId: t.ConfigId}
		err = mysql.InsertActUserKeepMonth(userKeepMonth)
		if err != nil {
			log.LogError(err.Error())
			return ErrorCode
		}
		fmt.Printf("userKeepMonthInitTaskStatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
		tmpTime = nextMonth
		keepMonth++
	}
	return SuccessCode

}

func userKeepMonthDailyTask(t model.Task) {
	//获取当前时间到6月前的时间列表
	currentTime, err := time.Parse("200601", time.Now().Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//计算的是上个月的数据
	currentTime = currentTime.AddDate(0, -1, 0)
	//遍历前6月的数据
	for i := 0; i < 7; i++ {
		startTime := currentTime.AddDate(0, -i, 0)
		//统计每月数据
		go func() {
			result := userKeepMonthDailyTaskStatistic(startTime, currentTime, t)
			if result == ErrorCode {
				mysql.FailRecord(startTime.Format("200601"), t.Id)

			}
		}()

	}
}

func userKeepMonthDailyTaskStatistic(startTime, currentTime time.Time, t model.Task) int {
	//计算keepMonth
	keepMonth := 0
	tmpTime := startTime
	for currentTime.After(tmpTime) {
		tmpTime = tmpTime.AddDate(0, 1, 0)
		keepMonth++
	}

	num, err := postgres.GetUserKeepCount(startTime, startTime.AddDate(0, 1, 0), currentTime, currentTime.AddDate(0, 1, 0), t)

	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	monthId, err := strconv.Atoi(startTime.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	userKeepMonth := &model.ActUserKeepMonth{MonthId: monthId, KeepMonth: keepMonth, Num: num, ConfigId: t.ConfigId}
	err = mysql.InsertActUserKeepMonth(userKeepMonth)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("userKeepMonthDailyTaskStatistic:fromday %v,currentTime:%v,num is:%v\n", startTime, currentTime, num)
	return SuccessCode
}
