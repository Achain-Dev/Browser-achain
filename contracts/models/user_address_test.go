package models

import (
	"testing"
	"fmt"
	"encoding/json"

)

func TestListByAddress(t *testing.T) {
	list, err := ListByAddress("ACT3khgGfz83sZNe5XM2moSBBu3TDVyNmHGc")

	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(list)
	fmt.Println(len(list))
	fmt.Println(string(bytes))


}
