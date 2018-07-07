package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
)

func InsertActUserKeepDay(userKeepDay *model.ActUserKeepDay) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserkeepday(dayId,keepDay,num,configId,addTime) values (?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(userKeepDay.DayId, userKeepDay.KeepDay, userKeepDay.Num, userKeepDay.ConfigId, time.Now(), userKeepDay.Num)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}

func InsertActUserKeepWeek(newUserWeek *model.NewUserWeek) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newuserweek (weekId,num,configId,addTime,startDay,endDay) values (?,?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(newUserWeek.WeekId, newUserWeek.Num, newUserWeek.ConfigId, time.Now(), newUserWeek.StartDay, newUserWeek.EndDay, newUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}

func InsertActUserKeepMonth(newUserMonth *model.NewUserMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newusermonth (monthId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newUserMonth.MonthId, newUserMonth.Num, newUserMonth.ConfigId, time.Now(), newUserMonth.Num)
	if err != nil {
		return err
	}
	return nil
}
