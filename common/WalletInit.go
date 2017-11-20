package common

import (
	"github.com/Masterminds/glide/path"
	"bytes"
	"github.com/robfig/config"
	"fmt"
)

var(
	WALLET_RPC,
	WALLET_NAME_PASSWORD string
)


func init()  {
	goPath := path.Gopath()
	var buffer bytes.Buffer
	buffer.WriteString(goPath)
	buffer.WriteString("/src/Browser-achain/conf/databaseConfig.ini")
	c, _ := config.ReadDefault(buffer.String())
	WALLET_RPC, _ = c.String("wallet", "wallet-rpc")
	WALLET_NAME_PASSWORD, _ = c.String("wallet", "wallet-name-password")
	fmt.Printf("wallet init | wallet_rpc=%s|wallet_name_password=%s\n",WALLET_RPC,WALLET_NAME_PASSWORD)

}
