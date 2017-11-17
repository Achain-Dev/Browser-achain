package models

import (
	"Browser-achain/common"
	"database/sql"
	"log"
)

type TbActTransaction struct {
	Id            int
	TrxId         *string
	BlockId       *string
	BlockNum      int64
	BlockPosition int
	TrxType       int
	CoinType      *string
	ContractId    *string
	FromAcct      *string
	FromAddr      *string
	ToAcct        *string
	ToAddr        *string
	SubAddress    *string
	Amount        *int64
	Fee           uint
	Memo          *string
	TrxTime       string
	CalledAbi     *string
	AbiParams     *string
	EventType     *string
	EventParam    *string
	ExtraTrxId    *string
	IsCompleted   uint8
	CreateTime    string
	UpdateTime    string
}

type ActTransactionDTO struct {
	Id int
	TrxId *string
	Amount string
	TrxType string
	CoinType *string
	TradeDescribe *string
	FromAddr *string
	BlockNum int64
	FromAcct *string
	ToAddr *string
	SubAddr *string
	ToAcct  *string
	CalledAbi *string
	TrxTime string
	IsCompleted string
	ContractId *string
	EventType *string
	EventParam *string
}

type ActTransactionVO struct {
	Data []ActTransactionDTO
	EndBlockNum int64

}


// Query an address transfer record and start block
func TransactionListQuery(start int64,userAddress,coinType string) ([]TbActTransaction,error) {

	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("TransactionListQuery|ERROR:", err)
		panic(err.Error())
	}

	var rows *sql.Rows
	if "" == coinType {
		stat, _ := db.Prepare("SELECT * FROM tb_act_transaction WHERE to_addr = ? AND block_num >= ? ORDER BY block_num ASC ")
		rows, err = stat.Query(userAddress,start)
	}else {
		stat, _ := db.Prepare("SELECT * FROM tb_act_transaction WHERE to_addr = ? AND block_num >= ? AND coin_type = ? ORDER BY block_num ASC")
		rows, err = stat.Query(userAddress,start,coinType)
	}

	if err != nil {
		log.Fatal("TransactionListQuery|ERROR:", err)
		panic(err.Error())
	}

	return mappingDataToTransactionList(rows)
}


func mappingDataToTransactionList(rows *sql.Rows) ([]TbActTransaction, error) {
	tbActTransactionList := make([]TbActTransaction, 0)
	for rows.Next() {
		var tbActTransaction TbActTransaction
		err := rows.Scan(
			&tbActTransaction.Id,
				&tbActTransaction.TrxId,
					&tbActTransaction.BlockId,
						&tbActTransaction.BlockNum,
							&tbActTransaction.BlockPosition,
								&tbActTransaction.TrxType,
									&tbActTransaction.CoinType,
										&tbActTransaction.ContractId,
											&tbActTransaction.FromAcct,
												&tbActTransaction.FromAddr,
													&tbActTransaction.ToAcct,
														&tbActTransaction.ToAddr,
															&tbActTransaction.SubAddress,
																&tbActTransaction.Amount,
																	&tbActTransaction.Fee,
																		&tbActTransaction.Memo,
																			&tbActTransaction.TrxTime,
																				&tbActTransaction.CalledAbi,
																					&tbActTransaction.AbiParams,
																						&tbActTransaction.EventType,
																							&tbActTransaction.EventParam,
																								&tbActTransaction.ExtraTrxId,
																									&tbActTransaction.IsCompleted,
																										&tbActTransaction.CreateTime,
																											&tbActTransaction.UpdateTime,

		)
		if err != nil {
			log.Fatal("mappingDataToTransactionList|ERROR:", err)
			panic(err)
			return tbActTransactionList, err
		}
		tbActTransactionList = append(tbActTransactionList, tbActTransaction)

	}
	return tbActTransactionList, nil
}
