package model

import "time"

type Funnel struct {
	FunnelId   int64
	FunnelName string
	Summary    string
	Status     int8
	AddTime    time.Time
	SystemId   int64
}

type FunnelData struct {
	Id       int64
	FunnelId int64
	DayId    int
	StepId   int64
	Num      int
	AddTime  time.Time
}

type FunnelStep struct {
	StepId    int64
	FunnelId  int64
	StepName  string
	Interfaces []string
	AddTime   time.Time
}
