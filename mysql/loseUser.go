package mysql

import (
	"codeboxUeba/model"
	"time"
	"codeboxUeba/log"
)

func InsertLoseUserWeek(loseUserWeek *model.LoseUserWeek) error {
	loseUserWeekSql := "insert into ueba_loseuserweek (weekId,num, startDay, endDay, addTime, configId) values (?,?,?,?,?,?) on duplicate key update num=?"
	stmt, err := db.Prepare(loseUserWeekSql)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(loseUserWeek.WeekId, loseUserWeek.Num, loseUserWeek.StartDay, loseUserWeek.EndDay, time.Now(), loseUserWeek.ConfigId, loseUserWeek.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func InsertLoseUserMonth(loseUserMonth *model.LoseUserMonth) error {
	backUserWeekSql := "insert into ueba_backusermonth (monthId,num, addTime, configId) values (?,?,?,?)  on duplicate key update num=?"
	stmt, err := db.Prepare(backUserWeekSql)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(loseUserMonth.MonthId, loseUserMonth.Num, time.Now(), loseUserMonth.ConfigId, loseUserMonth.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}
