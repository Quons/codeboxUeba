package mysql

import (
	"codeboxUeba/log"
	"codeboxUeba/model"
	"time"
)

func QueryFunnelList(systemId int) ([]int64, error) {
	funnelList := make([]int64, 0, 10)
	queryFunnelSql := "select funnelId from ueba_funnel WHERE systemId=? and status=1"
	stmt, err := db.Prepare(queryFunnelSql)
	if err != nil {
		log.LogError(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(systemId)
	if err != nil {
		log.LogError(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var fId int64
		rows.Scan(&fId)
		funnelList = append(funnelList, fId)
	}
	return funnelList, nil
}

func QueryFunnelSteps(fId int64) (step []*model.FunnelStep) {
	stepSql := "select stepId,funnelId,interfaces from ueba_funnelstep WHERE funnelId=?"
	stmt, err := db.Prepare(stepSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(fId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		funnelStep := &model.FunnelStep{}
		var interfaceIds string
		rows.Scan(&funnelStep.StepId, &funnelStep.FunnelId, &interfaceIds)
		funnelStep.Interfaces = QueryInterfacesById(interfaceIds)
		step = append(step, funnelStep)
	}
	return
}

func QueryInterfacesById(interfaceIds string) (urls []string) {
	interfaceIdSql := "select url from ueba_interface WHERE  interfaceId IN (" + interfaceIds + ")"
	rows, err := db.Query(interfaceIdSql)
	if err != nil {
		log.LogError(err.Error())
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var url string
		rows.Scan(&url)
		urls = append(urls, url)
	}
	return
}

func InsertFunnelData(funnelData *model.FunnelData) error {
	funnelDataSql := "INSERT INTO ueba_funneldata (funnelId, dayId, stepId, num, addTime) VALUE (?, ?, ?, ?, ?)  on duplicate key update num=?"
	stmt, err := db.Prepare(funnelDataSql)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(funnelData.FunnelId, funnelData.DayId, funnelData.StepId, funnelData.Num, time.Now(), funnelData.Num)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}
