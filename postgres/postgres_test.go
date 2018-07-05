package postgres

import (
	"testing"
)

func TestSqlSelect(t *testing.T) {
	sql:="SELECT * FROM dw_requestlog where platfrom=$1"
	dwSlice:=SqlSelect(sql,"ios")
	for _, value := range dwSlice {
		t.Logf("platform:%v\n",value)
	}
}
