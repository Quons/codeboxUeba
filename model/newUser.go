package model

import "time"

type NewUserDay struct {
	Id       int64
	DayId    int
	Num      int
	AddTime  time.Time
	ConfigId int64
}

type NewUserWeek struct {
	Id       int64
	WeekId   int
	Num      int
	StartDay int
	EndDay   int
	AddTime  time.Time
	ConfigId int64
}

type NewUserMonth struct {
	Id       int64
	MonthId    int
	Num      int
	AddTime  time.Time
	ConfigId int64
}

