package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"uebaDataJob/utils"
	"uebaDataJob/model"
	"uebaDataJob/log"
	"fmt"
	"strings"
)

var db *sql.DB

func Init() {
	var err error
	//db, err = sql.Open("mysql", conf.DB.Mysql.User+":"+conf.DB.Mysql.Pwd+"@tcp("+conf.DB.Mysql.Host+":"+strconv.Itoa(conf.DB.Mysql.Port)+")/"+conf.DB.Mysql.DbName+"?charset=utf8") //第一个参数为驱动名
	db, err = sql.Open("mysql", "test:123456@tcp(123.59.54.196:3333)/ueba?charset=utf8") //第一个参数为驱动名
	utils.CheckError(err)
}

func ReadConf(jobCode int) (conf []model.Task) {
	stmt, err := db.Prepare("select t.id,t.job_code,t.cursors,t.task_type,t.day_conf,t.week_conf,t.month_conf,dc.interfaces from task_conf t left join ueba_dataconfig dc on t.day_conf=dc.configId  where  job_code=?")
	utils.CheckError(err)
	rows, err := stmt.Query(jobCode)
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		task := model.Task{}
		rows.Scan(&task.Id, &task.JobCode, &task.Cursors, &task.TaskType, &task.DayConfigId,&task.WeekConfigId,&task.MonthConfigId, &task.Interface)
		if task.Interface == "" {
			log.LogError("interface is empty,please check you config")
			continue
		}
		conf = append(conf, task)
	}
	return
}

func UpdateCursor(task *model.Task) {
	stmt, err := db.Prepare("update task_conf set cursors=? where id=?")
	utils.CheckError(err)
	_, err = stmt.Exec(task.Cursors, task.Id)
	utils.CheckError(err)
}



func QueryInterfaceParamByConfig(configId int64) (interfaceParam string) {
	interfaceSql := "select interfaces from ueba_dataconfig where configId=?"
	stmt, err := db.Prepare(interfaceSql)
	utils.CheckError(err)
	rows, err := stmt.Query(configId)
	utils.CheckError(err)
	var interfaces string
	rows.Next()
	rows.Scan(&interfaces)
	if rows.Next() {
		log.LogError("too many result")
	}
	interfaceSql = "select url from ueba_interface where interfaceId in (" + interfaces + ")"
	rows, err = db.Query(interfaceSql)
	utils.CheckError(err)
	var interfaceSlice []string
	for rows.Next() {
		var i string
		rows.Scan(&i)
		interfaceSlice = append(interfaceSlice, i)
	}
	interfaceParam = fmt.Sprintf("'%s'", strings.Join(interfaceSlice, "','"))
	return
}

func FailRecord(date string, confId int) {
	recordSql := "update task_conf set fail_record = concat(ifnull(fail_record,''),?,',') where id = ?"
	stmt, err := db.Prepare(recordSql)
	if err!=nil{
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(date, confId)
	if err!=nil{
		log.LogError(err.Error())
		return
	}
}
