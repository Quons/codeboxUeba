package model

import "time"

type ActUserKeepDay struct {
	Id       int64
	DayId    int
	KeepDay  int
	Num      int
	AddTime  time.Time
	ConfigId int64
}

type ActUserKeepWeek struct {
	Id       int64
	WeekId   int
	KeepWeek int
	Num      int
	AddTime  time.Time
	ConfigId int64
}

type ActUserKeepMonth struct {
	Id       int64
	MonthId  int
	KeepWeek int
	Num      int
	AddTime  time.Time
	ConfigId int64
}
