package models

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	value, err := Get("go test 2")
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println(value)
	}
}

func TestSet(t *testing.T) {
	err := Set("go test 2", "go test value")

	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}

}

func TestSetWithExpire(t *testing.T)  {
	err := SetWithExpire("go test 2", "22", Redis_expire_time_EX, "3")
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
}
