package dto

import "strings"

var (
	WALLET_STATISTICS_ALL_TASK,
	WALLET_BLOCK_ALL_DATA,
	WALLET_TRANSACTION_ALL_DATA string

)

func GetWalletBlockRedisKey(uniqueKey string) string  {
	sl := []string{WALLET_BLOCK_ALL_DATA,uniqueKey}
	return strings.Join(sl,":")
}
