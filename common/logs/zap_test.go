package logs

import "testing"

func TestGetDailyLoggerName(t *testing.T) {
	name := getDailyLoggerName()
	t.Log(name)
}
