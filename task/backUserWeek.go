package task

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"strconv"
	"fmt"
	"codeboxUeba/postgres"
	"codeboxUeba/mysql"
)

//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据   ，要在活跃用户和新增用户任务跑完之后进行
func backUserWeekTask(t model.Task) {
	weekStatistic(t, backUserWeekStatistic)
}

func backUserWeekStatistic(t model.Task, fromDate time.Time, toDate time.Time) int {
	//查询当周的活跃用户
	year, week := fromDate.ISOWeek()
	WeekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}

	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//查询当周的留存用户 todo
	userKeepWeek, err := mysql.QueryUserKeepPreWeek(t.ConfigId, fromDate)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//查询当周的新增用户
	newUserWeek, err := mysql.QueryNewUserCurrentWeek(t.ConfigId, WeekId)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据
	backUserWeekCount := totalActUser - userKeepWeek - newUserWeek

	startDay, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	backUserWeek := &model.BackUserWeek{WeekId: WeekId, StartDay: startDay, EndDay: endDay, Num: backUserWeekCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	err = mysql.InsertBackUserWeek(backUserWeek)
	if err != nil {
		log.LogError(err.Error())
		return ErrorCode
	}
	fmt.Printf("backUserWeekStatistic:fromday %v,today %v, num is:%v\n", fromDate, toDate, backUserWeekCount)
	return SuccessCode
}
