package service

import (
	"Browser-achain/common"
	"Browser-achain/contracts/models"
	"Browser-achain/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type UserBalanceVo struct {
	CoinType string
	Balance  string
}

const coinType = "ACT"

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
		actualAmountString := util.GetActualAmount(value.Balance)
		actualAmount, _ := strconv.ParseFloat(actualAmountString, 32)

		// if coinType is not ACT and balance less than 0,ignore
		if coinType != userBalanceVo.CoinType && actualAmount <= float64(0) {
			return
		}
		// if coinType is ACT and balance less than 0, replace with 0
		if coinType == userBalanceVo.CoinType && actualAmount <= float64(0) {
			actualAmount = float64(0)
		}
		userBalanceVo.Balance = strconv.FormatFloat(actualAmount, 'E', -1, 32)
		userBalanceVoList = append(userBalanceVoList, userBalanceVo)
	}
	common.WebResultSuccess(userBalanceVoList, c)
}
