package models


import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListUrlsByContractId(t *testing.T) {
	tbActContractConfigList, err := ListUrlsByContractId("CON7w5yDZ5K4yxKjPwfn2seQjg8h6KLUwnCj")

	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(tbActContractConfigList)
	fmt.Println(len(tbActContractConfigList))
	fmt.Println(string(bytes))

}
