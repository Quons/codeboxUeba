package conf

import (
	"testing"
)

func TestInit(t *testing.T)  {
	//nit()
	for key, value := range Tasks {
		t.Log("key:",key,"  ","value:",value)
	}
	t.Log(DB.Mysql.Host)
}
