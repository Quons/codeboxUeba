package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"codeboxUeba/utils"
	"codeboxUeba/model"
	"codeboxUeba/log"
	"fmt"
	"strings"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("mysql", "test:123456@tcp(123.59.54.196:3333)/ueba?charset=utf8") //第一个参数为驱动名
	utils.CheckError(err)
}

func ReadConf(jobCode int) (conf []model.Task) {
	stmt, err := db.Prepare("SELECT t.id,t.job_code,t.cursors,t.task_type,t.config_id,dc.interfaces FROM task_conf t LEFT JOIN ueba_dataconfig dc ON t.config_id=dc.configId  WHERE  job_code=?")
	utils.CheckError(err)
	rows, err := stmt.Query(jobCode)
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		task := model.Task{}
		rows.Scan(&task.Id, &task.JobCode, &task.Cursors, &task.TaskType, &task.ConfigId, &task.Interface)
		if task.Interface == "" {
			log.LogError("interface is empty,please check you config")
			continue
		}
		conf = append(conf, task)
	}
	return
}

func UpdateCursor(task *model.Task) {
	stmt, err := db.Prepare("UPDATE task_conf SET cursors=? WHERE id=?")
	utils.CheckError(err)
	_, err = stmt.Exec(task.Cursors, task.Id)
	utils.CheckError(err)
}

func QueryInterfaceParamByConfig(configId int64) (interfaceParam string) {
	interfaceSql := "SELECT interfaces FROM ueba_dataconfig WHERE configId=?"
	stmt, err := db.Prepare(interfaceSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	rows, err := stmt.Query(configId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	var interfaces string
	rows.Next()
	rows.Scan(&interfaces)
	if rows.Next() {
		log.LogError("too many result")
		return
	}
	if interfaces == "" {
		log.LogError("interfaces is empty!")
		return
	}
	interfaceSql = "select url from ueba_interface where interfaceId in (" + interfaces + ")"
	rows, err = db.Query(interfaceSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}

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
	recordSql := "UPDATE task_conf SET fail_record = concat(ifnull(fail_record,''),?,',') WHERE id = ?"
	stmt, err := db.Prepare(recordSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	_, err = stmt.Exec(date, confId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}
