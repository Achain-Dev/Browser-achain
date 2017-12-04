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
		template := new(service.ActServiceTemplate)
		template.Browser = new(service.ActBrowserService)
		actBrowser.GET("/contract/balance/query/:address", template.Browser.QueryBalanceByAddress)
		actBrowser.GET("/contract/query/:page/:perPage", template.Browser.QueryContractByKey)
		actBrowser.GET("/getUserAddressBalance/:userAddress", template.Browser.QueryAddressInfo)
		actBrowser.GET("/transactionQuery/:userAddress/:start", template.Browser.TransactionListQuery)
		actBrowser.GET("/TransactionEx/Query/:page/:pageSize", template.Browser.TransactionExQuery)
		actBrowser.GET("/blockMaxNum/query", template.Browser.QueryBlockMaxNumber)
		actBrowser.GET("/block/query/:page/:pageSize", template.Browser.QueryBlockInfo)
		actBrowser.GET("/block/info/query", template.Browser.QueryBlockInfoByBlockIdOrNum)
		actBrowser.GET("/block/agent/query/:page/:pageSize", template.Browser.QueryBlockAgent)
		actBrowser.GET("/statistic/transaction", template.Browser.StatisticsTransaction)
	}

	//act wallet http
	actWallet := router.Group("/api/wallet/act")
	{
		broadcastTemplate := new(service.ActBroadcastTemplate)
		broadcastTemplate.Broadcast = new(service.ActBroadcastService)
		actWallet.POST("/network/broadcast/transaction",broadcastTemplate.Broadcast.NetworkBroadcastTransaction)
		actWallet.GET("/network/get/code",broadcastTemplate.Broadcast.NetworkGetCode)
		actWallet.POST("/network/broadcast/transactionWithCode",broadcastTemplate.Broadcast.NetworkBroadcastTransactionWithCode)

	}



	router.Run(":8381")
}
