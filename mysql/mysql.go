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
	stmt, err := db.Prepare("SELECT t.from_date,t.to_date, t.id,t.job_code,t.task_type,t.config_id,dc.interfaces FROM task_conf t LEFT JOIN ueba_dataconfig dc ON t.config_id=dc.configId  WHERE  job_code=?")
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(jobCode)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		task := model.Task{}
		rows.Scan(&task.FromDate, &task.ToDate, &task.Id, &task.JobCode,  &task.TaskType, &task.ConfigId, &task.Interface)
		if task.Interface == "" {
			log.LogError("interface is empty,please check you config")
			continue
		}
		conf = append(conf, task)
	}
	return
}

func UpdateCursor(task *model.Task) error {
	stmt, err := db.Prepare("UPDATE task_conf SET from_date='',to_date='' WHERE id=?")
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	defer stmt.Close()
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	_, err = stmt.Exec(task.Id)
	if err != nil {
		log.LogError(err.Error())
		return err
	}
	return nil
}

func QueryInterfaceParamByConfig(configId int64) (interfaceParam string, err error) {
	interfaceSql := "SELECT interfaces FROM ueba_dataconfig WHERE configId=?"
	stmt, err := db.Prepare(interfaceSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(configId)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	defer rows.Close()
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
	defer rows.Close()

	var interfaceSlice []string
	for rows.Next() {
		var i string
		rows.Scan(&i)
		interfaceSlice = append(interfaceSlice, i)
	}
	interfaceParam = fmt.Sprintf("'%s'", strings.Join(interfaceSlice, "','"))
	return
}

func RecordFail(failRecord *model.FailRecord) {
	failSql := "insert into task_fail_record (job_code, task_type, config_id,from_date,to_date) VALUE (?,?,?,?,?)"
	stmt, err := db.Prepare(failSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(failRecord.JobCode, failRecord.TaskType, failRecord.ConfigId, failRecord.FromDate, failRecord.ToDate)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}

func ReadFailRecord() (tasks []model.Task, err error) {
	failRecordSql := "select job_code,task_type,config_id,from_date,to_date from task_fail_record WHERE status=1"
	stmt, err := db.Prepare(failRecordSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		log.LogError(err.Error())
		return
	}
	for rows.Next() {
		var task model.Task
		rows.Scan(&task)
		tasks = append(tasks, task)
	}
	return
}

func CleanFailRecord() {
	cleanSql := "update task_fail_record SET status=0"
	_, err := db.Exec(cleanSql)
	if err != nil {
		log.LogError(err.Error())
		return
	}
}
