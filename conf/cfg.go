package conf

import (
	"codeboxUeba/model"
	"os"
	"io/ioutil"
	. "codeboxUeba/utils"
	"codeboxUeba/mysql"
	"encoding/json"
	"path/filepath"
)

var Tasks []model.Task
var DB *model.DB

func Init() {
	//数据库配置
	abs, err := filepath.Abs("")
	CheckError(err)
	DB = &model.DB{}
	err = json.Unmarshal(readFile(abs+"/db.json"), DB)
	CheckError(err)
	mysql.Init()
	Tasks = mysql.ReadConf(127)
}

func readFile(path string) []byte {
	cfgFile, err := os.Open(path)
	defer cfgFile.Close()
	CheckError(err)
	content, err := ioutil.ReadAll(cfgFile)
	CheckError(err)
	return content
}

const (
	ActUserDay   = "actUserDay"
	ActUserWeek  = "actUserWeek"
	ActUserMonth = "actUserMonth"
	NewUserDay   = "newUserDay"
	NewUserWeek  = "newUserWeek"
	NewUserMonth = "newUserMonth"
)
