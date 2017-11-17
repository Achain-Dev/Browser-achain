package models

import (
	"testing"
	"fmt"
)

func TestTransactionExQuery(t *testing.T) {
	list, err := TransactionExQuery("d9328c168d69af403300e8264428cf4a667f23e1", 1, 10)
	fmt.Println(list)
	fmt.Print(err)
}
