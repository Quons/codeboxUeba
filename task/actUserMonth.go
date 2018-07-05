package task

import (
	"sync"
	"uebaDataJob/model"
	"time"
	"uebaDataJob/postgres"
	"fmt"
	"strconv"
	"uebaDataJob/mysql"
	"uebaDataJob/log"
)

func actUserMonthTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//配置成-1 推出任务
	if t.MonthConfigId == -1 {
		wg.Done()
		return
	}
	//获取cursor到当前时间的时间列表
	fromDate, err := time.Parse("20060102", t.Cursors)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	// 获取本月的一号 todo 添加指定时间段的功能
	fromDate, err = time.Parse("200601", fromDate.Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	nowDate := time.Now()
	for {
		toDate := fromDate.AddDate(0, 1, 0)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}

		go actUserMonthStatistics(t, fromDate, toDate)
		//actUserDayChan <- &model.ActUserDayStatistic{t, fromDate, toDate}
		fromDate = toDate
	}
}

func actUserMonthStatistics(t model.Task, fromDate time.Time, toDate time.Time) {

	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[actUserMonth:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	num, err := postgres.GetGpCount(t.MonthConfigId, fromDate, toDate)
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
	actUserWeek := &model.ActUserMonth{Num: num, ConfigId: t.MonthConfigId, MonthId: monthId,}
	err = mysql.InsertActUserMonth(actUserWeek)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	fmt.Printf("fromday %v,today %v, num is:%v\n", fromDate, toDate, num)
}
