package mysql

import (
	"codeboxUeba/model"
	"codeboxUeba/utils"
	"time"
	"codeboxUeba/log"
)

func InsertNewUserDay(newUserDay *model.NewUserDay) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newuserday (dayId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	utils.CheckError(err)
	_, err = stmt.Exec(newUserDay.DayId, newUserDay.Num, newUserDay.ConfigId, time.Now(), newUserDay.Num)
	utils.CheckError(err)
}

func InsertNewUserWeek(newUserWeek *model.NewUserWeek) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newuserweek (weekId,num,configId,addTime,startDay,endDay) values (?,?,?,?,?,?) on duplicate key update num=?")
	if err!=nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(newUserWeek.WeekId, newUserWeek.Num, newUserWeek.ConfigId, time.Now(), newUserWeek.StartDay, newUserWeek.EndDay, newUserWeek.Num)
	if err!=nil {
		log.LogError(err.Error())
		return
	}
}

func InsertNewUserMonth(newUserMonth *model.NewUserMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newusermonth (monthId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err!=nil {
		return err
	}
	_, err = stmt.Exec(newUserMonth.MonthId, newUserMonth.Num, newUserMonth.ConfigId, time.Now(), newUserMonth.Num)
	if err!=nil {
		return err
	}
	return nil
}
