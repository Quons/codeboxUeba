package model

import "time"

type BackUserWeek struct {
	Id       int64
	WeekId   int
	StartDay int
	EndDay   int
	Num      int
	AddTime  time.Time
	ConfigId int64
}
type BackUserMonth struct {
	Id       int64
	MonthId   int
	Num      int
	AddTime  time.Time
	ConfigId int64
}
