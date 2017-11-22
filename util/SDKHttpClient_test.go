package util

import (
	"Browser-achain/common"
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	var params = []string{}
	//result := Post("http://10.23.1.198:18888/rpc", "admin:123456", "info", params)
	result := Post(common.WALLET_RPC, common.WALLET_NAME_PASSWORD, "info", params)
	fmt.Println(result)

}
