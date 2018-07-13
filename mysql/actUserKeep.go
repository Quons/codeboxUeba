package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
	"strconv"
	"fmt"
	"database/sql"
)

func InsertActUserKeepDay(userKeepDay *model.ActUserKeepDay) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("INSERT INTO ueba_actuserkeepday(dayId,keepDay,num,configId,addTime) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userKeepDay.DayId, userKeepDay.KeepDay, userKeepDay.Num, userKeepDay.ConfigId, time.Now(), userKeepDay.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertActUserKeepWeek(actUserKeepWeek *model.ActUserKeepWeek) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("insert into ueba_actuserkeepweek(weekId,keepWeek,num,configId,addTime) values (?,?,?,?,?) on duplicate key update num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(actUserKeepWeek.WeekId, actUserKeepWeek.KeepWeek, actUserKeepWeek.Num, actUserKeepWeek.ConfigId, time.Now(), actUserKeepWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertActUserKeepMonth(newUserMonth *model.ActUserKeepMonth) error {
	//检查dayid是否存在
	stmt, err := db.Prepare("INSERT INTO ueba_newusermonth (monthId,num,configId,addTime) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE num=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newUserMonth.MonthId, newUserMonth.Num, newUserMonth.ConfigId, time.Now(), newUserMonth.Num)
	if err != nil {
		return err
	}
	return nil
}

//获取上周的用户留存，用户流失和用户回流统计中使用
func QueryUserKeepPreWeek(confId int64, fromDate time.Time) (int, error) {
	fromDate = fromDate.AddDate(0, 0, -7)
	year, week := fromDate.ISOWeek()
	weekId, err := strconv.Atoi(fmt.Sprintf("%v%v", year, week))
	stmt, err := db.Prepare("SELECT num FROM ueba_actuserkeepweek WHERE weekId=? AND keepWeek=1 AND configId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	defer stmt.Close()
	var num = 0
	err = stmt.QueryRow(weekId, confId).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			log.LogError("QueryUserKeepPreWeek no result")
			return 0, nil
		} else {
			log.LogError(err.Error())
			return 0, err
		}
	}
	return num, nil
}

func QueryUserKeepPreMonth(confId int64, fromDate time.Time) (int, error) {
	fromDate = fromDate.AddDate(0, -1, 0)
	monthId := fromDate.Format("200601")
	stmt, err := db.Prepare("SELECT num FROM ueba_actuserkeepmonth WHERE monthId=? AND keepMonth=1 AND configId=?")
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	defer stmt.Close()
	var num = 0
	err = stmt.QueryRow(monthId, confId).Scan(&num)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		log.LogError(err.Error())
		return 0, err
	}
	return num, nil
}
