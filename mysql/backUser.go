package mysql

import (
	"codeboxUeba/model"
	"codeboxUeba/log"
	"time"
)

func InsertBackUserWeek(backUserWeek *model.BackUserWeek) error {
	backUserWeekSql := "insert into ueba_backuserweek (weekId,num, startDay, endDay, addTime, configId) values (?,?,?,?,?,?)  on duplicate key update num=?"
	stmt, err := db.Prepare(backUserWeekSql)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	_, err = stmt.Exec(backUserWeek.WeekId, backUserWeek.Num, backUserWeek.StartDay, backUserWeek.EndDay, time.Now(), backUserWeek.ConfigId, backUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertBackUserMonth(backUserMonth *model.BackUserMonth) error {
	backUserWeekSql := "insert into ueba_backusermonth (monthId,num, addTime, configId) values (?,?,?,?) on duplicate key update num=?"
	stmt, err := db.Prepare(backUserWeekSql)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	_, err = stmt.Exec(backUserMonth.MonthId, backUserMonth.Num, time.Now(), backUserMonth.ConfigId, backUserMonth.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}
