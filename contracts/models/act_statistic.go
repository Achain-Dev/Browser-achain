package models

import (
	"Browser-achain/contracts/dto"
	"encoding/json"
	"fmt"
)

type ActStatisticsDto struct {
	TransNum          uint64
	TransAmount       string
	AccountNum        uint64
	ContractNum       uint64
	TransactionHourly uint64
	TransactionPeak   uint64
}

func StatisticsAllDataForQuery() (ActStatisticsDto, error) {
	var actStatisticsDto ActStatisticsDto
	dtoString, err := Get(dto.WALLET_STATISTICS_ALL_TASK)
	if err != nil {
		return actStatisticsDto, err
	}
	err = json.Unmarshal([]byte(dtoString), &actStatisticsDto)
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	fmt.Println(actStatisticsDto)
	return actStatisticsDto, nil
}
