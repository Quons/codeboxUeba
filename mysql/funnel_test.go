package mysql

import "testing"

func TestQueryFunnelList(t *testing.T) {
	fIdList, err := QueryFunnelList(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(fIdList)

}

func TestQueryFunnelSteps(t *testing.T) {
	steps := QueryFunnelSteps(1)
	if len(steps) == 0 {
		t.Error("null steps")
		return
	}
	for _, step := range steps {
		t.Logf("%+v", step)
	}
}

func TestQueryInterfacesById(t *testing.T) {
	i := QueryInterfacesById("1,2,3")
	t.Log(i)
}
