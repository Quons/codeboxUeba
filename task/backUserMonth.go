package task

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/postgres"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
	"fmt"
)

func backUserMonthTask(t model.Task) {
	monthStatistic(t, backUserMonthStatistics)
}

func backUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) int {

	//查询当月的活跃用户
	monthId, err := strconv.Atoi(fromDate.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//查询当月的留存用户
	userKeepMonth, err := mysql.QueryUserKeepPreMonth(t.ConfigId, fromDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//查询当月的新增用户
	newUserMonth, err := mysql.QueryNewUserCurrentMonth(t.ConfigId, monthId)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据
	backUserMonthCount := totalActUser - userKeepMonth - newUserMonth

	backUserMonth := &model.BackUserMonth{MonthId: monthId, Num: backUserMonthCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	err = mysql.InsertBackUserMonth(backUserMonth)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("backUserMonthStatistics:fromday %v,today %v, num is:%v\n", fromDate, toDate, backUserMonthCount)
	return SuccessCode
}
