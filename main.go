package main

import (
	"codeboxUeba/task"
	"sync"
	"github.com/robfig/cron"
	"time"
	"codeboxUeba/model"
	"codeboxUeba/conf"
	"codeboxUeba/mysql"
)

var counter = 0

var resultChan = make(chan *model.Task)

var failTaskChan = make(chan *model.FailInfo)

func main() {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		run()
	})
	c.Start()
	for {
		select {
		case s := <-resultChan:
			//将数据更新到数据库中
			mysql.UpdateCursor(s)
		}
	}
}

func run() {
	//重新读取配置
	conf.Init()
	//copyConf := conf.Tasks
	//如果之前任务未完成，只运行当天的任务
	if counter != 0 {
		//获取当前时间的前一天
		nTime := time.Now()
		yesTime := nTime.AddDate(0, 0, -1)
		logDay := yesTime.Format("20060102")
		//修改配置中的cursor
		for i := 0; i < len(conf.Tasks); i++ {
			conf.Tasks[i].Cursors = logDay
		}
	}

	//计数器添加
	counter++
	wg := &sync.WaitGroup{}
	//获取任务列表
	for _, t := range conf.Tasks {
		t.FailChan = failTaskChan
		job := task.TasksFactory(t.TaskType)
		if job != nil {
			wg.Add(1)
			//todo 判断todate是否为null，为null就设置成当前时间
			go job(wg, resultChan, t)
		} else {
			continue
		}
	}
	//等待当前任务列表完成
	wg.Wait()
	//清空计数器
	counter--
}
