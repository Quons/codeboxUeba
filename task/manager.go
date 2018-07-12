package task

import (
	"codeboxUeba/model"
	. "codeboxUeba/conf"
	"time"
	"codeboxUeba/log"
	"codeboxUeba/mysql"
)

const (
	ErrorCode   = 0
	SuccessCode = 1
)

func TasksFactory(taskName string) (f func(task model.Task)) {
	switch taskName {
	case ActUserWeek:
		return actUserWeekTask
	case ActUserDay:
		return actUserDayTask
	case ActUserMonth:
		return actUserMonthTask
	case NewUserDay:
		return newUserDayTask
	case NewUserWeek:
		return newUserWeekTask
	case NewUserMonth:
		return newUserMonthTask
	case ActUserKeepDay:
		return actUserKeepDayTask
	case ActUserKeepWeek:
		return actUserKeepWeekTask
	case ActUserKeepMonth:
		return actUserKeepMonthTask
	case BackUserWeek:
		return backUserWeekTask
	case BackUserMonth:
		return backUserMonthTask
	case LoseUserWeek:
		return loseUserWeekTask
	case LoseUserMonth:
		return loseUserMonthTask
	case FunnelTask:
		return funnelTask
	default:
		return nil
	}
}

func dayStatistic(t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time) int) {
	//批量任务
	if t.FromDate != "" && t.ToDate != "" {
		fromDate, err := time.Parse("20060102", t.FromDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		if t.ToDate == "" {
			t.ToDate = time.Now().Format("20060102")
		}
		toDate, err := time.Parse("20060102", t.ToDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}

		tmpDate := toDate
		for fromDate.Before(toDate) {
			tmpDate = fromDate.AddDate(0, 0, 1)
			go func(from, to time.Time, t model.Task) {
				result := f(t, from, to)
				if result == ErrorCode {
					//记录错误信息
					mysql.FailRecord(from.Format("20060102"), t.Id)
				}
			}(fromDate, tmpDate, t)
			fromDate = tmpDate
		}
		//todo 修改fromdate todate
		mysql.UpdateCursor(&t)
	}

	//日常任务
	//获取任务区间
	toTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	fromTime := toTime.AddDate(0, 0, -1)
	go func() {
		result := f(t, fromTime, toTime)
		if result == ErrorCode {
			//记录错误信息
			mysql.FailRecord(fromTime.Format("20060102"), t.Id)
		}
	}()
}

func monthStatistic(t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time) int) {
	//批量任务
	if t.FromDate != "" && t.ToDate != "" {
		fromDate, err := time.Parse("200601", t.FromDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}

		if t.ToDate == "" {
			t.ToDate = time.Now().Format("200601")
		}
		toDate, err := time.Parse("200601", t.ToDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		tmpDate := toDate
		for fromDate.Before(toDate) {
			tmpDate = fromDate.AddDate(0, 1, 0)
			go func(from, to time.Time, t model.Task) {
				result := f(t, fromDate, tmpDate)
				if result == ErrorCode {
					//记录错误信息
					mysql.FailRecord(fromDate.Format("20060102"), t.Id)
				}
			}(fromDate, toDate, t)
			fromDate = tmpDate
		}
		//todo 修改fromdate todate
		mysql.UpdateCursor(&t)
	}

	//日常任务
	//获取任务区间
	toTime, err := time.Parse("200601", time.Now().Format("200601"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	fromTime := toTime.AddDate(0, -1, 0)
	go func() {
		result := f(t, fromTime, toTime)
		if result == ErrorCode {
			//记录错误信息
			mysql.FailRecord(fromTime.Format("200601"), t.Id)
		}
	}()
}

func weekStatistic(t model.Task, f func(t model.Task, fromDate time.Time, toDate time.Time) int) {
	//批量任务
	if t.FromDate != "" && t.ToDate != "" {
		fromDate, err := time.Parse("20060102", t.FromDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		//获取本周的周一
		fromDate = getMondayTime(fromDate)

		if t.ToDate == "" {
			t.ToDate = time.Now().Format("20060102")
		}

		toDate, err := time.Parse("20060102", t.ToDate)
		if err != nil {
			log.LogError(err.Error())
			return
		}
		//获取本周的周一
		toDate = getMondayTime(toDate)
		tmpDate := toDate
		for fromDate.Before(toDate) {
			tmpDate = fromDate.AddDate(0, 0, 7)
			go func(from, to time.Time, t model.Task) {
				result := f(t, fromDate, tmpDate)
				if result == ErrorCode {
					//记录错误信息
					mysql.FailRecord(fromDate.Format("20060102"), t.Id)
				}
			}(fromDate, toDate, t)
			fromDate = tmpDate
		}
		//todo 修改fromdate todate
		mysql.UpdateCursor(&t)
	}

	//日常任务
	//获取任务区间
	toTime, err := time.Parse("20060102", time.Now().Format("20060102"))
	if err != nil {
		log.LogError(err.Error())
		return
	}
	toTime = getMondayTime(toTime)

	fromTime := toTime.AddDate(0, 0, -7)
	go func() {
		result := f(t, fromTime, toTime)
		if result == ErrorCode {
			//记录错误信息
			mysql.FailRecord(fromTime.Format("200601"), t.Id)
		}
	}()
}

func getMondayTime(t time.Time) time.Time {
	//获取本周的周一
	toMonday := t.Weekday()
	for toMonday != time.Monday {
		t = t.AddDate(0, 0, -1)
		toMonday = t.Weekday()
	}
	return t
}
