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

func actUserDayTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	//配置成-1 推出任务
	if t.DayConfigId == -1 {
		wg.Done()
		return
	}
	//获取cursor到当前时间的时间列表 todo 添加指定时间段的功能
	fromDate, err := time.Parse("20060102", t.Cursors)
	nowDate := time.Now()
	utils.CheckError(err)
	for {
		toDate := fromDate.AddDate(0, 0, 1)
		//只执行到当前日期
		if toDate.After(nowDate) {
			//更新执行状态 数据库或者文件，暂时使用文件进行记录
			//使用channel进行通信
			t.Cursors = fromDate.Format("20060102")
			rc <- &t
			wg.Done()
			return
		}

		go statistics(t, fromDate, toDate)
		//actUserDayChan <- &model.ActUserDayStatistic{t, fromDate, toDate}
		fromDate = toDate
	}
}

func statistics(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord("[actUserDay:"+fromDate.Format("20060102")+"]", t.Id)
		}
	}()

	num, err := postgres.GetGpCount(t.DayConfigId, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		panic(err)
		return
	}

	//把查询到的数据插入到mysql中
	dayId, err := strconv.Atoi(fromDate.Format("20060102"))
	utils.CheckError(err)
	actUserDay := &model.ActUserDay{Num: num, ConfigId: t.DayConfigId, DayId: dayId}
	mysql.InsertActUserDay(actUserDay)
	fmt.Printf("fromday %v,num is:%v\n", fromDate, num)
}
