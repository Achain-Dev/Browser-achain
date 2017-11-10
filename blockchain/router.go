package blockchain

import (
	"Browser-achain/contracts/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	// act browser http
	actBrowser := router.Group("/api/browser/act")
	{
		actBrowser.GET("/contract/balance/query/:address", service.QueryBalanceByAddress)
	}

	//act wallet http
	router.Group("/api/wallet/act")
	{

	}

	router.Run(":8381")
}
