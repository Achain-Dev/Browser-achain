package service

import (
	"Browser-achain/contracts/models"
	"github.com/gin-gonic/gin"
	"Browser-achain/util"
	"log"
	"Browser-achain/common"
)

type UserBalanceVo struct {
	CoinType string
	Balance  string
}

const coinType  = "ACT"

// Check the balance of all currencies in the address according to the address
func QueryBalanceByAddress(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		common.WebResultFail(c)
	}

	list, err := models.ListByAddress(address)

	if err != nil {
		log.Fatal("QueryBalanceByAddress|query data ERROR:", err)
		common.WebResultFail(c)
	}

	userBalanceVoList := make([]UserBalanceVo, 0)

	for _, value := range list {
		var userBalanceVo UserBalanceVo
		userBalanceVo.CoinType = value.CoinType
		userBalanceVo.Balance = util.GetActualAmount(value.Balance)
		if coinType != userBalanceVo.CoinType && userBalanceVo.Balance == "0" {
			return
		}
		userBalanceVoList = append(userBalanceVoList, userBalanceVo)
	}
	common.WebResultSuccess(userBalanceVoList, c)
}
