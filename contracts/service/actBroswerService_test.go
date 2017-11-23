package service

import (
	"Browser-achain/contracts/models"
	"Browser-achain/util"
	"fmt"
	"encoding/json"
)

 func main()  {
	var tbActBlock  models.TbActBlock
	resultMap := make(map[string]interface{}, 0)

	tbActBlock.BlockId = "1"
	tbActBlock.BlockNum = uint64(2)
	tbActBlock.BlockSize = uint64(3)
	tbActBlock.Signee = "2213"
	tbActBlock.TransAmount = uint64(231231)
	tbActBlock.TransNum = uint(2312)
	tbActBlock.TransFee = uint64(321123)
	tbActBlock.BlockTime = "234123"
	if &tbActBlock != nil {
		resultMap["block_id"] = tbActBlock.BlockId
		resultMap["block_num"] = tbActBlock.BlockNum
		resultMap["block_size"] = tbActBlock.BlockSize
		resultMap["signee"] = tbActBlock.Signee
		resultMap["trans_amount"] = util.GetActualAmount(&int64(tbActBlock.TransAmount))
		resultMap["trans_num"] = tbActBlock.TransNum
		resultMap["block_bonus"] = 5
		resultMap["trans_fee"] = util.GetActualAmount(&int64(tbActBlock.TransFee))
		resultMap["block_time"] = tbActBlock.BlockTime
	}

	fmt.Println(json.Marshal(resultMap))
}
