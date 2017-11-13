package models

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestListContractInfoByKey(t *testing.T) {
	contractVOPage, _ := ListContractInfoByKey("SM", Forever, 1, 10, 1)
	fmt.Println(contractVOPage)
	bytes,_ := json.Marshal(contractVOPage)
	fmt.Println(string(bytes))
}
