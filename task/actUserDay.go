package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"codeboxUeba/utils"
	"fmt"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
)

func actUserDayTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	dayStatistic(wg, rc, t, actUserInsert)
}

func actUserInsert(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord(fromDate.Format("20060102"), t.Id)
		}
	}()

	num, err := postgres.GetGpCount(t.ConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	//把查询到的数据插入到mysql中
	dayId, err := strconv.Atoi(fromDate.Format("20060102"))
	utils.CheckError(err)
	actUserDay := &model.ActUserDay{Num: num, ConfigId: t.ConfigId, DayId: dayId}
	mysql.InsertActUserDay(actUserDay)
	fmt.Printf("actUserDay:fromday %v,num is:%v\n", fromDate, num)
}
