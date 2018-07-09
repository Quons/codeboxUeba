package mysql

import (
	"codeboxUeba/model"
	"qiniupkg.com/x/log.v7"
	"time"
)

func InsertBackUserWeek(backUserWeek *model.BackUserWeek) {
	backUserWeekSql := "insert into ueba_backuserweek (weekId,num, startDay, endDay, addTime, configId,) values (?,?,?,?,?,?)"
	stmt, err := db.Prepare(backUserWeekSql)
	if err != nil {
		log.Error(err.Error())
		return
	}
	_, err = stmt.Exec(backUserWeek.WeekId, backUserWeek.Num, backUserWeek.StartDay, backUserWeek.EndDay, time.Now(), backUserWeek.ConfigId)
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
