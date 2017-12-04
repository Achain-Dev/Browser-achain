package service

import (
	"Browser-achain/common"
	"Browser-achain/contracts/models"
	"Browser-achain/util"
	"Browser-achain/util/graph.verification"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"strings"
)

const CODE_PREFIX  = "act_broadcast_prefix"
type ActBroadcast interface {
	//  broadcast transaction
	NetworkBroadcastTransaction(c *gin.Context)
	// get capture verification
	NetworkGetCode(c *gin.Context)
	// broadcast transaction with code
	NetworkBroadcastTransactionWithCode(c *gin.Context)
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

func (_ *ActBroadcastService) NetworkGetCode(c *gin.Context) {

	codeKey := uuid.NewV4().String()
	code := graph_verification.GetRandomCode(4)
	imgBase64 := graph_verification.EncodeCodeToBash64(code)
	codeMap := make(map[string]string, 1)
	codeKeyRedis := strings.Join([]string{CODE_PREFIX, codeKey}, ":")
	codeMap["codeKey"] = codeKeyRedis
	codeMap["imgCodeKey"] = imgBase64
	fmt.Printf("NetworkGetCode|codeKey=%s|code=%s\n", codeKey, code)
	//300秒验证码过期
	models.SetWithExpire(codeKeyRedis, code, models.Redis_expire_time_EX, "300")
	common.WebResultSuccess(codeMap, c)
}


func (_ *ActBroadcastService) NetworkBroadcastTransactionWithCode(c *gin.Context) {
	message := c.PostForm("message")
	imgCodeKey := c.PostForm("imgCodeKey")
	code := c.PostForm("code")
	
	if code == "" || imgCodeKey == "" || message == "" {
		common.WebResultFail(c)
		return
	}
	codeKeyRedis := strings.Join([]string{CODE_PREFIX, imgCodeKey}, ":")
	value, err := models.Get(codeKeyRedis)

	if err != nil || value == "" {
		common.WebResultMiss(c,10003,"code expire")
		return
	}
	result := util.Post(common.WALLET_RPC, common.WALLET_NAME_PASSWORD, "network_broadcast_transaction", []string{message})
	common.WebResultSuccess(result, c)
}
