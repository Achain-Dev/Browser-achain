package util

import (
	"strconv"
	"time"
)

func GetTimeAddEight(date time.Time) time.Time {
	return date.Add(8 * time.Hour)
}

// calculate actual amount
func GetActualAmount(amount *int64) float64 {
	amountString := "0"
	if amount != nil {
		amountString = strconv.FormatInt(*amount, 64)
	}
	amountFloat, _ := strconv.ParseFloat(amountString, 64)
	actualAmount := amountFloat / float64(100000)
	return actualAmount
}
