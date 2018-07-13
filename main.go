package main

import (
	"codeboxUeba/task"
	"github.com/robfig/cron"
	"codeboxUeba/conf"
	"flag"
	"strings"
	"codeboxUeba/mysql"
	"codeboxUeba/log"
	"fmt"
)

func main() {
	//获取启动参数，startDate，endDate，如果startDate为空，则只执行日常定时任务，endDate默认为当前时间
	//jobCode 任务执行分组
	var jobCode int
	flag.IntVar(&jobCode, "jobCode", 1, "job group code")
	flag.Parse()
	//bundle job
	go run(jobCode, "Day")

	//daily job
	c := cron.New()
	specDay := "*/5 * * * * ?"
	c.AddFunc(specDay, func() {
		run(jobCode, "Day")
	})
	specWeek := "*/5 * * * * ?"
	c.AddFunc(specWeek, func() {
		run(jobCode, "Week")
	})
	specMonth := "*/5 * * * * ?"
	c.AddFunc(specMonth, func() {
		run(jobCode, "Month")
	})

	specFailRetry := "*/5 * * * * ?"
	c.AddFunc(specFailRetry, func() {
		reTry(jobCode)
	})
	c.Start()
	select {}
}

func reTry(jobCode int) {
	failRecords, err := mysql.ReadFailRecord()
	if err != nil {
		log.LogError(err.Error())
		return
	}

	for _, t := range failRecords {
		fmt.Println("cursor.......", t.FromDate)
		if t.JobCode == jobCode {
			job := task.TasksFactory(t.TaskType)
			if job != nil {
				//todo 判断todate是否为null，为null就设置成当前时间
				go job(t)
			} else {
				continue
			}
		}
	}
	//执行完之后清除失败记录
	mysql.CleanFailRecord()

}

func run(jobCode int, taskType string) {
	//重新读取配置
	conf.Init()
	//获取任务列表
	for _, t := range conf.Tasks {
		//fmt.Println("cursor.......",t.Cursors)
		if t.JobCode == jobCode && strings.Contains(t.TaskType, taskType) {
			job := task.TasksFactory(t.TaskType)
			if job != nil {
				//todo 判断todate是否为null，为null就设置成当前时间
				go job(t)
			} else {
				continue
			}
		}
	}

}
