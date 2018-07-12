package task

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/postgres"
	"fmt"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
)

func actUserWeekTask(t model.Task) {
	weekStatistic(t, actUserWeekStatistics)
}

func actUserWeekStatistics(t model.Task, fromDate time.Time, toDate time.Time) int {

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//把查询到的数据插入到mysql中
	weekId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	actUserWeek := &model.ActUserWeek{Num: num, ConfigId: t.ConfigId, WeekId: weekId, StartDay: weekId, EndDay: endDay}
	err = mysql.InsertActUserWeek(actUserWeek)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("actUserWeek:fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
	return SuccessCode
}
