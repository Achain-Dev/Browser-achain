package util

import (
	"testing"
	"time"
	"fmt"
)

func TestGetTimeAddEight(t *testing.T) {
	now := time.Now
	fmt.Println(GetTimeAddEight(now()))
}
