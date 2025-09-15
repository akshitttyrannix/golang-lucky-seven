package functions

import "time"

func GetStartOfDay() int64 {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return startOfDay.Unix()
}

func GetEndOfDay() int64 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return endOfDay.Unix()
}

func GetToday() int64 {
	now := time.Now()
	return int64(now.Day())
}

func GetDateFormatted() string {
	now := time.Now()
	return now.Format("020106")
}

func CurrentTime() int64 {
	now := time.Now()
	return int64(now.Unix())
}
