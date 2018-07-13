package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"database/sql"
)

func InsertNewUserDay(newUserDay *model.NewUserDay) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("INSERT INTO ueba_newuserday (dayId,num,configId,addTime) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newUserDay.DayId, newUserDay.Num, newUserDay.ConfigId, time.Now(), newUserDay.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertNewUserWeek(newUserWeek *model.NewUserWeek) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("INSERT INTO ueba_newuserweek (weekId,num,configId,addTime,startDay,endDay) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newUserWeek.WeekId, newUserWeek.Num, newUserWeek.ConfigId, time.Now(), newUserWeek.StartDay, newUserWeek.EndDay, newUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertNewUserMonth(newUserMonth *model.NewUserMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("INSERT INTO ueba_newusermonth (monthId,num,configId,addTime) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE num=?")
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
	stmt, err := db.Prepare("SELECT num FROM ueba_newuserweek WHERE configId=? AND weekId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	num := 0
	err = stmt.QueryRow(confId, weekId).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.LogError(err.Error())
		return 0, err
	}
	return num, nil
}
func QueryNewUserCurrentMonth(confId int64, weekId int) (int, error) {
	stmt, err := db.Prepare("SELECT num FROM ueba_newusermonth WHERE configId=? AND monthId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	defer stmt.Close()
	num := 0
	err = stmt.QueryRow(confId, weekId).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.LogError(err.Error())
		return 0, err
	}
	return num, nil
}
