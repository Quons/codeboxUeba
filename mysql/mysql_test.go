package mysql

import (
	"testing"
	"codeboxUeba/model"
)

func init() {
	Init()

}
func TestQueryInterfaceByConfig(t *testing.T) {
	r, err := QueryInterfaceParamByConfig(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}

func TestQueryInterfaceParamByConfig(t *testing.T) {
	s, err := QueryInterfaceParamByConfig(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
}

func TestRecordFail(t *testing.T) {
	r := &model.FailRecord{JobCode: 123, TaskType: "actUserDay", ConfigId: 1, FromDate: "20060102", ToDate: "20070102"}
	RecordFail(r)
}
