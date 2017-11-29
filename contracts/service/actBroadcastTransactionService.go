package service

import (
	"Browser-achain/common"
	"Browser-achain/util"
	"github.com/gin-gonic/gin"
)

type ActBroadcast interface {
	//  broadcast transaction
	NetworkBroadcastTransaction(c *gin.Context)
}

type ActBroadcastTemplate struct {
	Broadcast ActBroadcast
}

type ActBroadcastService struct {
}

// post request method x-www-form-urlencoded
func (_ *ActBroadcastService) NetworkBroadcastTransaction(c *gin.Context) {
	message := c.PostForm("message")

	if message == "" {
		common.WebResultFail(c)
		return
	}

	result := util.Post(common.WALLET_RPC, common.WALLET_NAME_PASSWORD, "network_broadcast_transaction", []string{message})
	common.WebResultSuccess(result, c)
}
