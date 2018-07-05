package model

type Task struct{
	Id int
	JobCode int
	Cursors string
	TaskType string
	DayConfigId int64
	WeekConfigId int64
	MonthConfigId int64
	Interface string
	FailChan chan *FailInfo
}

type DB struct {
	Mysql DBConf  `json:"mysql"`
	Postgres DBConf `json:"postgres"`
}

type DBConf struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Pwd string `json:"pwd"`
	User string `json:"user"`
	DbName string `json:"dbName"`
}


type FailInfo struct {
	date string
	taskId int
}