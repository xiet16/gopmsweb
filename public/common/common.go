package common

import (
	"time"
)

//查找某值是否在数组中
func InArrayString(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

func StrToTime(dates string) int64 {
	tm2, _ := time.Parse("2006-01-02", dates)
	return tm2.Unix()
}

func StrToDatetime(dates string) int64 {
	tm2, _ := time.Parse("2006-01-02 15:04:05", dates)
	return tm2.Unix()
}
