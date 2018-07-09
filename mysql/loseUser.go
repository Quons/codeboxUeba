package mysql

import (
	"codeboxUeba/model"
	"qiniupkg.com/x/log.v7"
	"time"
)

func InsertLoseUserWeek(loseUserWeek *model.LoseUserWeek) {
	loseUserWeekSql := "insert into ueba_loseuserweek (weekId,num, startDay, endDay, addTime, configId,) values (?,?,?,?,?,?)"
	stmt, err := db.Prepare(loseUserWeekSql)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_, err = stmt.Exec(loseUserWeek.WeekId, loseUserWeek.Num, loseUserWeek.StartDay, loseUserWeek.EndDay, time.Now(), loseUserWeek.ConfigId)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
func InsertBackUserMonth(backUserMonth *model.BackUserMonth) {
	backUserWeekSql := "insert into ueba_backusermonth (weekId,num, addTime, configId) values (?,?,?,?)"
	stmt, err := db.Prepare(backUserWeekSql)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_, err = stmt.Exec(backUserMonth.MonthId, backUserMonth.Num, time.Now(), backUserMonth.ConfigId)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
