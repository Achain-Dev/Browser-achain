package util

import (
	"time"
	"strconv"
)

func GetTimeAddEight(date time.Time) time.Time {
	return date.Add(8 * time.Hour)
}

// calculate actual amount
func GetActualAmount(amount *int64) string {

	if amount == nil || *amount == int64(0) {
		return "0"
	}
	actualAmount := float64(*amount) / float64(100000)
	return strconv.FormatFloat(actualAmount, 'f', -1, 32)
}
