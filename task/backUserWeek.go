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

func backUserWeekStatistic(t model.Task, fromDate time.Time, toDate time.Time) {
	//查询当周的活跃用户
	year, week := fromDate.ISOWeek()
	WeekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}

	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	//查询当周的留存用户 todo
	userKeepWeek := 0
	for userKeepWeek == 0 {
		time.Sleep(5 * time.Minute)
		log.LogInfo("backUserWeek waiting for userKeepWeek")
		userKeepWeek, err = mysql.QueryUserKeepPreWeek(t.ConfigId, fromDate)
		if err != nil {
			log.LogError(err.Error())
			RecordFailTask(fromDate, toDate, &t)
			continue
		}
	}

	//查询当周的新增用户
	newUserWeek := 0
	for newUserWeek == 0 {
		time.Sleep(5 * time.Minute)
		log.LogInfo("backUserWeek waiting for newUserWeek")
		newUserWeek, err = mysql.QueryNewUserCurrentWeek(t.ConfigId, WeekId)
		if err != nil {
			log.LogError(err.Error())
			RecordFailTask(fromDate, toDate, &t)
			continue
		}
	}

	//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据
	backUserWeekCount := totalActUser - userKeepWeek - newUserWeek

	startDay, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	backUserWeek := &model.BackUserWeek{WeekId: WeekId, StartDay: startDay, EndDay: endDay, Num: backUserWeekCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	err = mysql.InsertBackUserWeek(backUserWeek)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	fmt.Printf("backUserWeekStatistic:fromday %v,today %v, num is:%v\n", fromDate, toDate, backUserWeekCount)
}
