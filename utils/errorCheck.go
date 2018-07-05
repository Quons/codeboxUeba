package utils

import (
	"os"
	"uebaDataJob/log"
	"fmt"
)

// 错误检查这边好好写下
func CheckErrorDetail(err error,msg string,date interface{})  {
	if err!=nil {
		log.LogError(fmt.Sprintf("err:[%s] message:[%s] date:[%v]",err.Error(),msg,date))
		os.Exit(1)
	}
}

func CheckError(err error)  {
	if err!=nil {
		log.LogError(err.Error())
		os.Exit(1)
	}
}
