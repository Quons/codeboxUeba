package task

import (
	"sync"
	"uebaDataJob/model"
	"time"
	"uebaDataJob/utils"
	"uebaDataJob/postgres"
	"fmt"
	"strconv"
	"uebaDataJob/mysql"
	"uebaDataJob/log"
)

func actUserWeekTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//配置成-1 推出任务
	if t.WeekConfigId == -1 {
		wg.Done()
		return
	}
	//获取cursor到当前时间的时间列表 todo 添加指定时间段的功能
	fromDate, err := time.Parse("20060102", t.Cursors)
	utils.CheckError(err)

	//获取本周的周一
	weekDay := fromDate.Weekday()
	for weekDay != time.Monday {
		fromDate = fromDate.AddDate(0, 0, -1)
		weekDay = fromDate.Weekday()
	}

	nowDate := time.Now()
	for {
		toDate := fromDate.AddDate(0, 0, 7)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}

		go actUserWeekStatistics(t, fromDate, toDate)
		//actUserDayChan <- &model.ActUserDayStatistic{t, fromDate, toDate}
		fromDate = toDate
	}
}

func actUserWeekStatistics(t model.Task, fromDate time.Time, toDate time.Time) {

	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[actUserWeek:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	num ,err:= postgres.GetGpCount(t.WeekConfigId, fromDate, toDate)
	if err!=nil {
		log.LogError(err.Error())
		panic(err)
		return
	}
	//把查询到的数据插入到mysql中
	weekId, err := strconv.Atoi(fromDate.Format("20060102"))
	utils.CheckError(err)
	endDay, err := strconv.Atoi(toDate.Format("20060102"))
	utils.CheckError(err)
	actUserWeek := &model.ActUserWeek{Num: num, ConfigId: t.WeekConfigId, WeekId: weekId, StartDay: weekId, EndDay: endDay}
	mysql.InsertActUserWeek(actUserWeek)
	fmt.Printf("fromday %v,today %v, num is:%v\n", fromDate, toDate,num)
}
