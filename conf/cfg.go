package conf

import (
	"codeboxUeba/model"
	"os"
	"io/ioutil"
	. "codeboxUeba/utils"
	"codeboxUeba/mysql"
	"encoding/json"
	"path/filepath"
	"codeboxUeba/log"
)

var Tasks []model.Task
var DB *model.DB

func Init() {
	//数据库配置
	abs, err := filepath.Abs("")
	if err != nil {
		log.LogError(err.Error())
		os.Exit(1)
	}
	DB = &model.DB{}
	err = json.Unmarshal(readFile(abs+"/db.json"), DB)
	if err != nil {
		log.LogError(err.Error())
		os.Exit(1)
	}
	mysql.Init()
	Tasks = mysql.ReadConf()
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
	ActUserDay       = "actUserDay"
	ActUserWeek      = "actUserWeek"
	ActUserMonth     = "actUserMonth"
	NewUserDay       = "newUserDay"
	NewUserWeek      = "newUserWeek"
	NewUserMonth     = "newUserMonth"
	ActUserKeepDay   = "actUserKeepDay"
	ActUserKeepWeek  = "actUserKeepWeek"
	ActUserKeepMonth = "actUserKeepMonth"
	BackUserWeek     = "backUserWeek"
	BackUserMonth    = "backUserMonth"
	LoseUserWeek     = "loseUserWeek"
	LoseUserMonth    = "loseUserMonth"
	FunnelTask       = "funnelTaskDay"
)
