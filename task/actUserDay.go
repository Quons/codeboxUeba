package task

import (
	"codeboxUeba/model"
	"time"
	"fmt"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
)

func actUserDayTask(t model.Task) {
	dayStatistic(t, ActUserDayInsert)
}

func ActUserDayInsert(t model.Task, fromDate time.Time, toDate time.Time) {
	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	//把查询到的数据插入到mysql中
	dayId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	actUserDay := &model.ActUserDay{Num: num, ConfigId: t.ConfigId, DayId: dayId}
	err = mysql.InsertActUserDay(actUserDay)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	fmt.Printf("actUserDay:fromday %v,num is:%v\n", fromDate, num)
}
