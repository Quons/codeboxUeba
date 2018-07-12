package model

import "time"

type LoseUserWeek struct {
	Id       int64
	WeekId   int
	StartDay int
	EndDay   int
	Num      int
	AddTime  time.Time
	ConfigId int64
}
type LoseUserMonth struct {
	Id       int64
	MonthId  int
	Num      int
	AddTime  time.Time
	ConfigId int64
}
