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

func newUserMonthTask(t model.Task) {
	monthStatistic(t, newUserMonthStatistics)
}

func newUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) {

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	//把查询到的数据插入到mysql中
	monthId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	newUserWeek := &model.NewUserMonth{Num: num, ConfigId: t.ConfigId, MonthId: monthId,}
	err = mysql.InsertNewUserMonth(newUserWeek)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	fmt.Printf("newUserMonth:fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
}
