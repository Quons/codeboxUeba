package model

import "time"

type ActUserDay struct {
	Id       int64
	DayId    int
	Num      int
	AddTime  time.Time
	ConfigId int64
}

type ActUserWeek struct {
	Id       int64
	WeekId   int
	Num      int
	StartDay int
	EndDay   int
	AddTime  time.Time
	ConfigId int64
}

type ActUserMonth struct {
	Id       int64
	MonthId    int
	Num      int
	AddTime  time.Time
	ConfigId int64
}


type ActUserDayStatistic struct {
	TaskConf Task
	FromDate time.Time
	ToDate   time.Time
}
