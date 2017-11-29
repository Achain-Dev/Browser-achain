package service

import (
	"Browser-achain/common"
	"Browser-achain/contracts/models"
	"Browser-achain/util"
	"github.com/gin-gonic/gin"
	"github.com/grafana/grafana/pkg/components/simplejson"
	"log"
	"strconv"
	"strings"
)

type ActService interface {
	// Check the balance of all currencies in the address according to the address
	QueryBalanceByAddress(c *gin.Context)
	// Check the address balance by the keyword
	QueryContractByKey(c *gin.Context)
	// Query the act address account information
	QueryAddressInfo(c *gin.Context)
	// Query transactions by address and block number
	TransactionListQuery(c *gin.Context)
	// Query by origin trx id
	TransactionExQuery(c *gin.Context)
	// Block max number query
	QueryBlockMaxNumber(c *gin.Context)
	// Query block info
	QueryBlockInfo(c *gin.Context)
	// Query block info by block id ,or block number
	QueryBlockInfoByBlockIdOrNum(c *gin.Context)
	// Query block info by signee
	QueryBlockAgent(c *gin.Context)
	// Statistical home related information
	StatisticsTransaction(c *gin.Context)
}

type ActServiceTemplate struct {
	Browser ActService
}

type ActBrowserService struct {

}

type UserBalanceVo struct {
	CoinType string
	Balance  string
}

const coinType = "ACT"
const CONTRACT_PREFIX = "CON"

func (_ *ActBrowserService) QueryBalanceByAddress(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		common.WebResultFail(c)
		return
	}

	list, err := models.ListByAddress(address)

	if err != nil {
		log.Fatal("QueryBalanceByAddress|query data ERROR:", err)
		common.WebResultFail(c)
		return
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

func (_ *ActBrowserService) QueryContractByKey(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("perPage"))
	keyword := c.Query("keyword")
	log.Printf("QueryContractByKey|page=%s|perPage=%s|keyword=%s\n", page, perPage, keyword)
	if page < 1 || perPage < 1 {
		common.WebResultFail(c)
		return
	}

	queryType := 1
	// keyword is not empty and keyword startWith CONN and the length of keyword greater than 30
	if keyword != "" && strings.Index(keyword, CONTRACT_PREFIX) == 0 && len(keyword) > 30 {
		queryType = 0
	}
	contractInfoPageVO, err := models.ListContractInfoByKey(keyword, models.Forever, page, perPage, queryType)
	if err != nil {
		common.WebResultFail(c)
		return
	}

	contractInfoVOList := make([]models.ContractInfoVO, 0)
	tbActContractInfoList := contractInfoPageVO.ActContractInfoList
	for _, actContractInfo := range tbActContractInfoList {
		var contractInfoVO models.ContractInfoVO
		circulation := actContractInfo.Circulation
		intCirculation, _ := strconv.ParseInt(util.GetActualAmount(&circulation), 10, 64)
		contractInfoVO.Circulation = intCirculation
		contractInfoVO.CoinType = actContractInfo.CoinType
		contractInfoVO.CoinAddress = actContractInfo.OwnerAddress
		contractInfoVO.ContractName = actContractInfo.Name
		contractInfoVO.RegisterTime = actContractInfo.RegTime
		contractInfoVO.Status = actContractInfo.Status
		contractInfoVO.ContractId = actContractInfo.ContractId
		contractInfoVO.Coin = int(actContractInfo.Type)

		contractInfoVOList = append(contractInfoVOList, contractInfoVO)
	}
	contractInfoPageVO.ActContractInfoList = nil
	contractInfoPageVO.ContractInfoVOList = contractInfoVOList
	common.WebResultSuccess(contractInfoPageVO, c)
}

func (_ *ActBrowserService) QueryAddressInfo(c *gin.Context) {
	userActAddress := c.Param("userAddress")
	userAddressList, err := models.ListByAddressAndCoinType(userActAddress, "ACT")
	if err != nil {
		common.WebResultFail(c)
		return
	}
	var userAddressVO models.UserAddressVO

	if len(userAddressList) > 0 {
		tbUserAddress := userAddressList[0]
		actualAmount := util.GetActualAmount(tbUserAddress.Balance)
		userAddressVO.Balance = actualAmount
		userAddressVO.Address = *tbUserAddress.UserAddress
	}
	common.WebResultSuccess(userAddressVO, c)
}

func (_ *ActBrowserService) TransactionListQuery(c *gin.Context) {
	userActAddress := c.Param("userAddress")
	start, _ := strconv.ParseInt(c.Param("start"), 10, 64)

	tbActTransactionList, err := models.TransactionListQuery(start, userActAddress, "ACT")
	if err != nil {
		common.WebResultFail(c)
		return
	}

	if len(tbActTransactionList) == 0 {
		common.WebResultMiss(c, 10002, "no more transactions")
		return
	}

	actTransactionDTOList := make([]models.ActTransactionDTO, 0)

	for _, tbActTransaction := range tbActTransactionList {
		var actTransactionDTO models.ActTransactionDTO
		actTransactionDTO.Id = tbActTransaction.Id
		actTransactionDTO.TrxId = tbActTransaction.TrxId
		actTransactionDTO.FromAcct = tbActTransaction.FromAcct
		actTransactionDTO.FromAddr = tbActTransaction.FromAddr
		actTransactionDTO.ToAcct = tbActTransaction.ToAcct
		actTransactionDTO.ToAddr = tbActTransaction.ToAddr
		actTransactionDTO.CalledAbi = tbActTransaction.CalledAbi
		actTransactionDTO.TrxType = strconv.Itoa(tbActTransaction.TrxType)
		actTransactionDTO.Amount = util.GetActualAmount(tbActTransaction.Amount)
		actTransactionDTO.TrxTime = tbActTransaction.TrxTime
		actTransactionDTO.IsCompleted = strconv.FormatUint(uint64(tbActTransaction.IsCompleted), 10)
		actTransactionDTO.SubAddr = tbActTransaction.SubAddress
		actTransactionDTO.CoinType = tbActTransaction.CoinType
		actTransactionDTO.BlockNum = tbActTransaction.BlockNum
		actTransactionDTOList = append(actTransactionDTOList, actTransactionDTO)
	}
	data := make(map[string]interface{}, 4)
	data["data"] = actTransactionDTOList
	data["code"] = 200
	data["msg"] = "success"
	data["endBlockNum"] = actTransactionDTOList[len(actTransactionDTOList)-1].BlockNum
	common.WebResultSuccessWithMap(c, data)
}

func (_ *ActBrowserService) TransactionExQuery(c *gin.Context) {
	originTrxId := c.DefaultQuery("originTrxId", "")
	page, _ := strconv.Atoi(c.Param("page"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	list, err := models.TransactionExQuery(originTrxId, page, pageSize)
	if err != nil {
		common.WebResultFail(c)
		return
	}
	common.WebResultSuccess(list, c)
}

func (_ *ActBrowserService) QueryBlockMaxNumber(c *gin.Context) {
	var params = []string{}
	result := util.Post(common.WALLET_RPC, common.WALLET_NAME_PASSWORD, "blockchain_get_block_count", params)
	var blockNum = int64(0)
	if result != "" {
		resultJson, err := simplejson.NewJson([]byte(result))
		if err != nil {
			panic(err.Error())
			log.Fatal("QueryBlockMaxNumber|getBlockNum|error convert json string")
			common.WebResultFail(c)
			return
		}
		blockNum = resultJson.Get("result").MustInt64()
	}
	common.WebResultSuccess(blockNum, c)
}

func (_ *ActBrowserService) QueryBlockInfo(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	actBlockPageVO, err := models.BlockQueryByPage("", page, pageSize)
	if err != nil {
		common.WebResultFail(c)
		return
	}
	common.WebResultSuccess(actBlockPageVO, c)
}

func (_ *ActBrowserService) QueryBlockInfoByBlockIdOrNum(c *gin.Context) {
	blockId := c.DefaultQuery("blockId", "")
	blockNum := c.DefaultQuery("blockNum", "")

	if blockId == "" && blockNum == "" {
		common.WebResultMiss(c, 10007, "param missing")
		return
	}

	var tbActBlock models.TbActBlock
	var err error
	if blockId != "" {
		result, errTemp := models.BlockQueryByBlockId(blockId)
		tbActBlock = *result
		err = errTemp
	} else {
		num, _ := strconv.ParseInt(blockNum, 10, 64)
		result, errTemp := models.BlockQueryByBlockNum(num)
		tbActBlock = *result
		err = errTemp
	}

	if err != nil {
		common.WebResultFail(c)
		return
	}
	resultMap := make(map[string]interface{}, 0)

	if &tbActBlock != nil {
		resultMap["block_id"] = tbActBlock.BlockId
		resultMap["block_num"] = tbActBlock.BlockNum
		resultMap["block_size"] = tbActBlock.BlockSize
		resultMap["signee"] = tbActBlock.Signee
		amount := int64(tbActBlock.TransAmount)
		resultMap["trans_amount"] = util.GetActualAmount(&amount)
		resultMap["trans_num"] = tbActBlock.TransNum
		resultMap["block_bonus"] = 5
		fee := int64(tbActBlock.TransFee)
		resultMap["trans_fee"] = util.GetActualAmount(&fee)
		resultMap["block_time"] = tbActBlock.BlockTime
	}
	common.WebResultSuccess(resultMap, c)
}

func (_ *ActBrowserService) QueryBlockAgent(c *gin.Context) {

	signee := c.DefaultQuery("signee", "")
	page, _ := strconv.Atoi(c.Param("page"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))

	if page < 1 || pageSize < 1 {
		common.WebResultFail(c)
		return
	}
	actBlockPageVO, err := models.BlockQueryByPage(signee, page, pageSize)
	if err != nil {
		common.WebResultFail(c)
		return
	}

	common.WebResultSuccess(actBlockPageVO, c)
}

func (_ *ActBrowserService) StatisticsTransaction(c *gin.Context) {
	actStatisticsDto, err := models.StatisticsAllDataForQuery()
	if err != nil {
		common.WebResultFail(c)
		return
	}
	common.WebResultSuccess(actStatisticsDto, c)
}
