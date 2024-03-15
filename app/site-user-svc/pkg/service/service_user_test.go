package service

import (
	"testing"
	"time"
)

func TestCheckZeroValue(t *testing.T) {
	a1 := "123"
	a2 := ""
	b1 := 234
	b2 := 0
	c1 := ""
	c2 := "xyz"
	d1 := time.Now()
	var d2 time.Time
	t.Log(checkZeroValue(a1, a2).(string))
	t.Log(checkZeroValue(b1, b2).(int))
	t.Log(checkZeroValue(c1, c2).(string))
	t.Log(checkZeroValue(d1, d2).(time.Time))
}
