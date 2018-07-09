package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"codeboxUeba/utils"
	"codeboxUeba/postgres"
	"fmt"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
)

func newUserWeekTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	weekStatistic(wg, rc, t, newUserWeekStatistics)
}

func newUserWeekStatistics(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[actUserWeek:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		panic(err)
		return
	}
	//把查询到的数据插入到mysql中 todo 修改weekid
	weekId, err := strconv.Atoi(fromDate.Format("20060102"))
	utils.CheckError(err)
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	utils.CheckError(err)
	newUserWeek := &model.NewUserWeek{Num: num, ConfigId: t.ConfigId, WeekId: weekId, StartDay: weekId, EndDay: endDay}
	mysql.InsertNewUserWeek(newUserWeek)
	fmt.Printf("newUserWeek:fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
}
