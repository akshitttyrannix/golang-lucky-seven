package functions

import "time"

func GetStartOfDay() uint32 {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return uint32(startOfDay.Unix())
}

func GetEndOfDay() uint32 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return uint32(endOfDay.Unix())
}

func GetToday() uint32 {
	now := time.Now()
	return uint32(now.Day())
}

func GetDateFormatted() string {
	now := time.Now()
	return now.Format("020106")
}

func CurrentTime() uint32 {
	now := time.Now()
	return uint32(now.Unix())
}
