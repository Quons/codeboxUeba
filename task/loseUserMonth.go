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

func loseUserMonthTask(t model.Task) {
	monthStatistic(t, loseUserMonthStatistics)
}

func loseUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) {

	//查询上个月的活跃用户
	monthId, err := strconv.Atoi(fromDate.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate.AddDate(0, -1, 0), toDate.AddDate(0, -1, 0))
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	//查询当月的留存用户 todo 暂时这么处理
	userKeepMonth := 0
	for userKeepMonth == 0 {
		time.Sleep(5 * time.Minute)
		log.LogInfo("loseUserMonth waiting for userKeepMonth")
		userKeepMonth, err = mysql.QueryUserKeepPreMonth(t.ConfigId, fromDate)
		if err != nil {
			log.LogError(err.Error())
			continue
		}
	}
	newUserMonth := 0
	for newUserMonth == 0 {
		time.Sleep(5 * time.Minute)
		log.LogInfo("loseUserMonth waiting for newUserMonth")
		//查询当月的新增用户
		newUserMonth, err = mysql.QueryNewUserCurrentMonth(t.ConfigId, monthId)
		if err != nil {
			log.LogError(err.Error())
			continue
		}
	}

	//月流失用户= 上月活跃用户-当月留存用户-当月新增用户  只需要查询上月数据
	loseUserMonthCount := totalActUser - userKeepMonth - newUserMonth

	loseUserMonth := &model.LoseUserMonth{MonthId: monthId, Num: loseUserMonthCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	err = mysql.InsertLoseUserMonth(loseUserMonth)
	if err != nil {
		log.LogError(err.Error())
		RecordFailTask(fromDate, toDate, &t)
		return
	}
	fmt.Printf("loseUserMonthStatistics:fromday %v,today %v, num is:%+v\n", fromDate, toDate, loseUserMonthCount)
}
