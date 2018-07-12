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

func newUserDayTask(t model.Task) {
	dayStatistic(t, newUserDayInsert)
}

func newUserDayInsert(t model.Task, fromDate time.Time, toDate time.Time) int {
	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//把查询到的数据插入到mysql中
	dayId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	newUserDay := &model.NewUserDay{Num: num, ConfigId: t.ConfigId, DayId: dayId}
	err = mysql.InsertNewUserDay(newUserDay)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("new userDay ,fromday %v,num is:%v\n", fromDate, num)
	return SuccessCode
}
