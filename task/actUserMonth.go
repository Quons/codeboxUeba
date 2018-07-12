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

func actUserMonthTask(t model.Task) {
	monthStatistic(t, actUserMonthStatistics)
}

func actUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) int {

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//把查询到的数据插入到mysql中
	monthId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	actUserWeek := &model.ActUserMonth{Num: num, ConfigId: t.ConfigId, MonthId: monthId,}
	err = mysql.InsertActUserMonth(actUserWeek)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("actUserMonth:fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
	return SuccessCode
}
