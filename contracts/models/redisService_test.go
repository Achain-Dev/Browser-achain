package models

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	value, err := Get("WALLET_BLOCK_ALL_DATA：")
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println(value)
	}
}

func TestSet(t *testing.T) {
	err := Set("go test", "go test value")

	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}

}
