package log

import (
	"github.com/op/go-logging"
	"os"
	"fmt"
	"time"
	"bytes"
	"runtime"
)

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

var today = "2006-01-02"

// 判断文件夹是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func fileInit() {
	fileName := "logs/log_"+time.Now().Format("2006-01-02")+".txt"
	exist,err := FileExists(fileName)
	if !exist {
		logfile,_ := os.Create(fileName)
		logfile.Close()
	}
	logFile, err := os.OpenFile(fileName, os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)
}

func LogInfo(message string)  {
	pc,file,_,_ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	log.Info(fmt.Sprintf("%s	%s",GetLogPrefix(file, f.Name(), "Info"), message))
	//fmt.Println(message)
}

func LogWarning(message string)  {
	pc,file,_,_ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	log.Warning(fmt.Sprintf("%s	%s",GetLogPrefix(file, f.Name(), "Warning"), message))
}

func LogError(message string)  {
	pc,file,_,_ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	log.Error(fmt.Sprintf("%s	%s",GetLogPrefix(file, f.Name(), "Error"), message))
}

func GetLogPrefix(file string, funcName string, level string) string {
	if time.Now().Format("2006-01-02") != today {
		fileInit()
		today = time.Now().Format("2006-01-02")
	}
	var buf bytes.Buffer
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	buf.WriteString("	[")
	buf.WriteString(level)
	buf.WriteString("]	")
	buf.WriteString(file)
	buf.WriteString("(")
	buf.WriteString(funcName)
	buf.WriteString("):")
	return buf.String()
}