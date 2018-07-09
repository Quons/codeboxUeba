package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"strconv"
	"fmt"
	"codeboxUeba/postgres"
	"codeboxUeba/mysql"
)

//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据   ，要在活跃用户和新增用户任务跑完之后进行
func loseUserWeekTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	weekStatistic(wg, rc, t, loseUserWeekStatistic)
	wg.Done()
}

func loseUserWeekStatistic(t model.Task, fromDate time.Time, toDate time.Time) {
	//查询上周的活跃用户
	year, week := fromDate.ISOWeek()
	weekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	if err != nil {
		log.LogError(err.Error())
		return
	}

	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate.AddDate(0, 0, -7), toDate.AddDate(0, 0, -7))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//查询当周的留存用户
	userKeepWeek, err := mysql.QueryUserKeepPreWeek(t.ConfigId, fromDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//查询当周的新增用户
	newUserWeek, err := mysql.QueryNewUserCurrentWeek(t.ConfigId, weekId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//周回流用户= 上周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据
	loseUserWeekCount := totalActUser - userKeepWeek - newUserWeek

	startDay, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	loseUserWeek := &model.LoseUserWeek{WeekId: weekId, StartDay: startDay, EndDay: endDay, Num: loseUserWeekCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	mysql.InsertLoseUserWeek(loseUserWeek)
}
