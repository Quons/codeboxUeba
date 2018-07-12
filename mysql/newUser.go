package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"errors"
)

func InsertNewUserDay(newUserDay *model.NewUserDay) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newuserday (dayId,num,configId,addTime) values (?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	_, err = stmt.Exec(newUserDay.DayId, newUserDay.Num, newUserDay.ConfigId, time.Now(), newUserDay.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertNewUserWeek(newUserWeek *model.NewUserWeek) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_newuserweek (weekId,num,configId,addTime,startDay,endDay) values (?,?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	_, err = stmt.Exec(newUserWeek.WeekId, newUserWeek.Num, newUserWeek.ConfigId, time.Now(), newUserWeek.StartDay, newUserWeek.EndDay, newUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertNewUserMonth(newUserMonth *model.NewUserMonth) error {
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

//获取当前周的新增用户
func QueryNewUserCurrentWeek(confId int64, weekId int) (int, error) {
	stmt, err := db.Prepare("select num from ueba_newuserweek where configId=? and weekId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}

	rows, err := stmt.Query(confId, weekId)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	num := 0
	rows.Next()
	rows.Scan(&num)
	if rows.Next() {
		log.LogError("too manny result")
		return 0, errors.New("too many result")
	}
	return num, nil
}
func QueryNewUserCurrentMonth(confId int64, weekId int) (int, error) {
	stmt, err := db.Prepare("select num from ueba_newusermonth where configId=? and monthId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}

	rows, err := stmt.Query(confId, weekId)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	num := 0
	rows.Next()
	rows.Scan(&num)
	if rows.Next() {
		log.LogError("too manny result")
		return 0, errors.New("too many result")
	}
	return num, nil
}
