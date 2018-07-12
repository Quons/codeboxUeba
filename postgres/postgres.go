package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"codeboxUeba/model"
	"codeboxUeba/log"
	"codeboxUeba/mysql"
	"time"
	"strings"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gpadmin:gpadmin@117.50.2.54:5433/ueba?sslmode=disable")
	checkErr(err)
}

func GetGpCount(confId int64, fromDate time.Time, toDate time.Time) (num int, err error) {
	//获取日活接口列表，进行预处理
	interfaceParam := mysql.QueryInterfaceParamByConfig(confId)
	if interfaceParam == "" {
		log.LogError("interfaceParam is empty")
		return 0, err
	}
	//查询，插入操作。 日活：统计的是userid的数量，接口可能有多个
	countSql := `select count(distinct userid)
			from dw_requestlog
			where
  				requesturl in (` + interfaceParam + `) and
				requesttime between $1 and $2
			`
	num, err = QueryCount(countSql, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	return
}

func FunnelCount(interfaces []string, fromDate time.Time, toDate time.Time) (num int, err error) {
	//查询，插入操作。 日活：统计的是userid的数量，接口可能有多个
	countSql := `select count(distinct userid)
			from dw_requestlog
			where
  				requesturl in ('` + strings.Join(interfaces, "','") + `') and
				requesttime between $1 and $2
			`
	num, err = QueryCount(countSql, fromDate, toDate)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	return

}

func GetUserKeepCount(startTime, startTimeD, endTime, endTimeD time.Time, t model.Task) (num int, err error) {
	//获取日活接口列表，进行预处理
	interfaceParam := mysql.QueryInterfaceParamByConfig(t.ConfigId)
	if interfaceParam == "" {
		log.LogError("interfaceParam is empty")
		return
	}
	userKeepSql := `
		select count(1) as count
		from (select distinct userid
      		from dw_requestlog
      		where
        		requesturl in (` + interfaceParam + `) and requesttime between $1 and $2
     		) d1 left join (select  distinct userid
                     		from dw_requestlog
                     		where requesturl in (` + interfaceParam + `) and
                           		requesttime between $3 and $4
                    		) d2 on d1.userid = d2.userid
					`
	num, err = QueryCount(userKeepSql, startTime, startTimeD, endTime, endTimeD)
	if err != nil {
		log.LogError(err.Error())
		return 0, err
	}
	return
}

func SqlSelect(sql string, params ...interface{}) []*model.Postgres {
	//查询数据
	stat, err := db.Prepare(sql)
	//rows, err := db.Query("SELECT * FROM dw_requestlog")
	checkErr(err)
	rows, err := stat.Query(params...)
	checkErr(err)
	defer rows.Close()
	dwSlice := make([]*model.Postgres, 0, 0)
	for rows.Next() {
		dw := &model.Postgres{}
		err = rows.Scan(&dw.Logid, &dw.Requesturl, &dw.Postdata, &dw.Getdata, &dw.Hostip, &dw.Userid, &dw.Statuscode,
			&dw.Responsecode, &dw.Responsedata, &dw.Spendtime, &dw.Requesttime, &dw.Source, &dw.Platform, &dw.Version, &dw.Systemid, &dw.Addtime)
		checkErr(err)
		//fmt.Println("sss")
		dwSlice = append(dwSlice, dw)
	}
	return dwSlice
}

func QueryCount(sql string, params ...interface{}) (num int, err error) {
	//查询数据
	stat, err := db.Prepare(sql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	checkErr(err)
	rows, err := stat.Query(params...)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&num)
	if rows.Next() {
		log.LogError("too many result")
		return
	}
	return
}

func CloseGp() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
}
