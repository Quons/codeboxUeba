package mysql

import (
	"testing"
)

func TestQueryInterfaceByConfig(t *testing.T) {
	Init()
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
