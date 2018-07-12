package task

import (
	"testing"
	"codeboxUeba/model"
	"time"
)

func TestFunnelInsert(t *testing.T) {
	funnelInsert(model.Task{Id: 1, ConfigId: 1}, time.Now().AddDate(0, 0, -1), time.Now())
}
