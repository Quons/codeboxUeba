package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"codeboxUeba/model"
	"codeboxUeba/log"
	"codeboxUeba/mysql"
	"time"
	"strings"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gpadmin:gpadmin@117.50.2.54:5433/ueba?sslmode=disable")
	if err != nil {
		log.LogError(err.Error())
		os.Exit(1)
	}
}

func GetGpCount(confId int64, fromDate time.Time, toDate time.Time) (num int, err error) {
	//获取日活接口列表，进行预处理
	interfaceParam, err := mysql.QueryInterfaceParamByConfig(confId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
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
	interfaceParam, err := mysql.QueryInterfaceParamByConfig(t.ConfigId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
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

func QueryCount(sql string, params ...interface{}) (num int, err error) {
	//查询数据
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(params...).Scan(&num)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	return
}
