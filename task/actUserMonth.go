package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"codeboxUeba/postgres"
	"fmt"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
)

func actUserMonthTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	monthStatistic(wg, rc, t, actUserMonthStatistics)
}

func actUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[actUserMonth:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//把查询到的数据插入到mysql中
	monthId, err := strconv.Atoi(fromDate.Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	actUserWeek := &model.ActUserMonth{Num: num, ConfigId: t.ConfigId, MonthId: monthId,}
	err = mysql.InsertActUserMonth(actUserWeek)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	fmt.Printf("actUserMonth:fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
}
