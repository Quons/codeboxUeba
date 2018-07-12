package task

import (
	"sync"
	"codeboxUeba/model"
	"time"
	"strconv"
	"codeboxUeba/mysql"
	"codeboxUeba/postgres"
	"codeboxUeba/log"
	"fmt"
)

func funnelTask(wg *sync.WaitGroup, rc chan *model.Task, t model.Task) {
	dayStatistic(wg, rc, t, funnelInsert)
}

func funnelInsert(t model.Task, fromDate time.Time, toDate time.Time) {
	defer func() {
		if recover() != nil {
			//如果失败，记录失败记录
			mysql.FailRecord(fromDate.Format("20060102"), t.Id)
		}
	}()
	//todo sytemId暂时写死
	//获取funnel 列表，遍历
	fIdList, err := mysql.QueryFunnelList(1)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	for _, fId := range fIdList {
		//获取flunnelid对应的step的interface
		funnelSteps := mysql.QueryFunnelSteps(fId)
		//遍历steps，分别进行查询
		for _, step := range funnelSteps {
			//查询接口的访问量

			num, err := postgres.FunnelCount(step.Interfaces, fromDate, toDate)
			if err != nil {
				log.LogError(err.Error())
				return
			}

			//插入到data表中
			dayId, err := strconv.Atoi(fromDate.Format("20060102"))
			if err != nil {
				log.LogError(err.Error())
				return
			}
			funnelData := &model.FunnelData{Num: num, FunnelId: fId, StepId: step.StepId, DayId: dayId}
			fmt.Printf("funnelInsert:%+v\n", funnelData)
			err = mysql.InsertFunnelData(funnelData)
			if err != nil {
				log.LogError(err.Error())
				continue
			}
		}
	}
}
