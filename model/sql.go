package model

import "time"

type Postgres struct {
	Logid int
	Requesturl string
	Postdata string
	Getdata string
	Hostip string
	Userid  string
	Statuscode int
	Responsecode string
	Responsedata string
	Spendtime int
	Requesttime time.Time
	Source string
	Platform string
	Version string
	Systemid int
	Addtime  time.Time
}
