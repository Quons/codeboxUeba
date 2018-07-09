package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"codeboxUeba/postgres"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
)

func backUserMonthTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	monthStatistic(wg, rc, t, backUserMonthStatistics)
}

func backUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[backUserMonth:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	//查询当月的活跃用户
	monthId, err := strconv.Atoi(fromDate.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	totalActUser, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//查询当月的留存用户
	userKeepMonth, err := mysql.QueryUserKeepPreMonth(t.ConfigId, fromDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//查询当月的新增用户
	newUserMonth, err := mysql.QueryNewUserCurrentMonth(t.ConfigId, monthId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//周回流用户= 当周活跃用户-当周留存用户-当周新增用户  只需要查询上周数据
	backUserMonthCount := totalActUser - userKeepMonth - newUserMonth

	backUserMonth := &model.BackUserMonth{MonthId: monthId, Num: backUserMonthCount, ConfigId: t.ConfigId}
	//将结果添加到表中
	mysql.InsertBackUserMonth(backUserMonth)
}
