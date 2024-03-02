package tools

import (
	"errors"
	"site/protocol/shared"
	"time"
)

// TimeToInt64 日期格式转换时间戳
func TimeToInt64(t time.Time) (int64, error) {
	if t.IsZero() {
		return 0, errors.New(shared.CodeMessageIgnoreCode(shared.TimeInvalid))
	}
	return t.Unix(), nil
}

// Int64ToTime 时间戳转换日期格式
func Int64ToTime(t int64) (time.Time, error) {
	if t <= 0 {
		return time.Time{}, errors.New(shared.CodeMessageIgnoreCode(shared.TimeInvalid))
	}
	return time.Unix(t, 0), nil
}

// CompareInt32Int 比较 int32 与 int 大小
func CompareInt32Int(a int32, b int) bool {
	if int(a) != b {
		return false
	}
	return true
}
