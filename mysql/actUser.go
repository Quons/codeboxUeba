package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
)

func InsertActUserDay(actUserDay *model.ActUserDay) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserday (dayId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(actUserDay.DayId, actUserDay.Num, actUserDay.ConfigId, time.Now(), actUserDay.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertActUserWeek(actUserWeek *model.ActUserWeek) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserweek (weekId,num,configId,addTime,startDay,endDay) values (?,?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(actUserWeek.WeekId, actUserWeek.Num, actUserWeek.ConfigId, time.Now(), actUserWeek.StartDay, actUserWeek.EndDay, actUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertActUserMonth(actUserMonth *model.ActUserMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actusermonth (monthId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(actUserMonth.MonthId, actUserMonth.Num, actUserMonth.ConfigId, time.Now(), actUserMonth.Num)
	if err != nil {
		return err
	}
	return nil
}
