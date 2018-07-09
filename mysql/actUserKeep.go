package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"errors"
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

func InsertActUserKeepWeek(actUserKeepWeek *model.ActUserKeepWeek) {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserkeepweek(weekId,keepWeek,num,configId,addTime) values (?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(actUserKeepWeek.WeekId, actUserKeepWeek.KeepWeek, actUserKeepWeek.Num, actUserKeepWeek.ConfigId, time.Now(), actUserKeepWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}

func InsertActUserKeepMonth(newUserMonth *model.ActUserKeepMonth) error {
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

func QueryUserKeepPreWeek(confId int64, weekId int) (int, error) {

	weekId = weekId - 1
	stmt, err := db.Prepare("select num from ueba_actuserkeepweek where weekId=? and keepWeek=1 and configId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	rows, err := stmt.Query(weekId, confId)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	var num = 0
	rows.Next()
	rows.Scan(&num)
	if rows.Next() {
		log.LogError("too many result!")
		err = errors.New("too many result")
		return 0, err
	}
	return num, nil

}
