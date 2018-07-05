package conf

import (
	"uebaDataJob/model"
	"os"
	"io/ioutil"
	."uebaDataJob/utils"
	"uebaDataJob/mysql"
	"encoding/json"
	"path/filepath"
)

var Tasks []model.Task
var DB *model.DB

func Init() {
	//数据库配置
	abs,err:=filepath.Abs("")
	CheckError(err)
	DB=&model.DB{}
	err=json.Unmarshal(readFile(abs+"/db.json"),DB)
	CheckError(err)
	mysql.Init()
	Tasks=mysql.ReadConf(127)
}

func readFile(path string)[]byte  {
	cfgFile,err:=os.Open(path)
	defer cfgFile.Close()
	CheckError(err)
	content,err:=ioutil.ReadAll(cfgFile)
	CheckError(err)
	return content
}

const (
	ActUser="actUser"
	NewUser="newUser"
)
