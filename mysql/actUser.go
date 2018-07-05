package mysql

import (
	"uebaDataJob/model"
	"uebaDataJob/utils"
	"time"
	"uebaDataJob/log"
)

func InsertActUserDay(actUserDay *model.ActUserDay) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserday (dayId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	utils.CheckError(err)
	_, err = stmt.Exec(actUserDay.DayId, actUserDay.Num, actUserDay.ConfigId, time.Now(), actUserDay.Num)
	utils.CheckError(err)
}

func InsertActUserWeek(actUserWeek *model.ActUserWeek) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserweek (weekId,num,configId,addTime,startDay,endDay) values (?,?,?,?,?,?) on duplicate key update num=?")
	if err!=nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(actUserWeek.WeekId, actUserWeek.Num, actUserWeek.ConfigId, time.Now(), actUserWeek.StartDay, actUserWeek.EndDay, actUserWeek.Num)
	if err!=nil {
		log.LogError(err.Error())
		return
	}
}

func InsertActUserMonth(actUserMonth *model.ActUserMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actusermonth (monthId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err!=nil {
		return err
	}
	_, err = stmt.Exec(actUserMonth.MonthId, actUserMonth.Num, actUserMonth.ConfigId, time.Now(), actUserMonth.Num)
	if err!=nil {
		return err
	}
	return nil
}
