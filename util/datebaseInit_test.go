package util

import (
	"testing"
	"errors"
	"fmt"

)

func TestGetDbConnection(t *testing.T) {
	db, err := GetDbConnection()

	if err != nil{
		errors.New("get  connection fail")
	}

	fmt.Println(db)
}

