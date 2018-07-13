package mysql

import (
	"testing"
	"codeboxUeba/model"
)

func init() {
	Init()

}
func TestQueryInterfaceByConfig(t *testing.T) {
	r := QueryInterfaceParamByConfig(1)
	t.Log(r)
}

func TestFailRecord(t *testing.T) {
	Init()
	FailRecord("20180705", 5)
}

func TestQueryInterfaceParamByConfig(t *testing.T) {
	s := QueryInterfaceParamByConfig(1)
	t.Log(s)
}

func TestRecordFail(t *testing.T) {
	r := &model.FailRecord{JobCode: 123, TaskType: "actUserDay", ConfigId: 1, FromDate: "20060102", ToDate: "20070102"}
	RecordFail(r)
}
