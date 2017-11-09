package util

import "time"

func GetTimeAddEight(date time.Time) (time.Time) {
	return date.Add(8*time.Hour)
}
