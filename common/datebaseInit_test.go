package common

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetDbConnection(t *testing.T) {
	db, err := GetDbConnection()

	if err != nil {
		errors.New("get  connection fail")
	}

	fmt.Println(db)
}
