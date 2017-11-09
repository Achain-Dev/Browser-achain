package util

import (
	"testing"
	"fmt"
)

func TestPost(t *testing.T) {
	var params = []string{}
	result := Post("http://172.16.33.201:18888/rpc", "admin:123456", "info", params)
	fmt.Println(result)
}
