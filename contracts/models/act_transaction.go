package models

import (
	"Browser-achain/common"
	"database/sql"
	"log"
	"strings"
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
	Id            int
	TrxId         *string
	Amount        string
	TrxType       string
	CoinType      *string
	TradeDescribe *string
	FromAddr      *string
	BlockNum      int64
	FromAcct      *string
	ToAddr        *string
	SubAddr       *string
	ToAcct        *string
	CalledAbi     *string
	TrxTime       string
	IsCompleted   string
	ContractId    *string
	EventType     *string
	EventParam    *string
}

type ActTransactionPage struct {
	Data         []TbActTransaction `json:"data"`
	CurrentPage  uint               `json:"currentPage"`
	PageSize     uint               `json:"pageSize"`
	TotalPage    uint               `json:"totalPage"`
	TotalRecords uint               `json:"totalRecords"`
}

type ActTransactionVO struct {
	Data        []ActTransactionDTO
	EndBlockNum int64
}

// Query an address transfer record and start block
func TransactionListQuery(start int64, userAddress, coinType string) ([]TbActTransaction, error) {

	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("TransactionListQuery|ERROR:", err)
		panic(err.Error())
	}

	var rows *sql.Rows
	if "" == coinType {
		stat, _ := db.Prepare("SELECT * FROM tb_act_transaction WHERE to_addr = ? AND block_num >= ? ORDER BY block_num ASC ")
		rows, err = stat.Query(userAddress, start)
	} else {
		stat, _ := db.Prepare("SELECT * FROM tb_act_transaction WHERE to_addr = ? AND block_num >= ? AND coin_type = ? ORDER BY block_num ASC")
		rows, err = stat.Query(userAddress, start, coinType)
	}

	if err != nil {
		log.Fatal("TransactionListQuery|ERROR:", err)
		panic(err.Error())
	}

	return mappingDataToTransactionList(rows)
}

func TransactionQueryByTrxId(trxId string) (*TbActTransaction, error) {
	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("TransactionQueryByTrxId|ERROR:", err)
		panic(err.Error())
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM tb_act_transaction WHERE trx_id = ? LIMIT 1", trxId)

	if err != nil {
		log.Fatal("TransactionQueryByTrxId|ERROR:", err)
		panic(err.Error())
		return nil, err
	}

	tbTransactionList, _ := mappingDataToTransactionList(rows)

	if len(tbTransactionList) == 0 {
		return nil, nil
	}
	return &tbTransactionList[0], nil

}

// Query transaction by
func TransactionListQueryByBlock(blockNum uint64, acctAddress string, page, pageSize int) (ActTransactionPage, error) {
	db, err := common.GetDbConnection()
	var actTransactionPage ActTransactionPage
	defer db.Close()

	if err != nil {
		log.Fatal("TransactionListQueryByBlock|ERROR:", err)
		panic(err.Error())
		return actTransactionPage, err
	}

	rows, countRows, err := getRowsAndCountRows(blockNum, acctAddress, err, db, page, pageSize)

	if err != nil {
		log.Fatal("TransactionListQueryByBlock|ERROR:", err)
		panic(err.Error())
		return actTransactionPage, err
	}

	totalRecords := countNumber(countRows)

	if totalRecords == 0 {
		return actTransactionPage, nil
	}

	totalPage := totalRecords/pageSize + 1
	if totalRecords%pageSize == 0 {
		totalPage = totalRecords / pageSize
	}

	actTransactionPage.TotalPage = uint(totalPage)
	actTransactionPage.CurrentPage = uint(page)
	actTransactionPage.PageSize = uint(pageSize)
	actTransactionPage.TotalRecords = uint(totalRecords)
	tbTransactionList, _ := mappingDataToTransactionList(rows)
	actTransactionPage.Data = tbTransactionList
	return actTransactionPage, nil

}

func getRowsAndCountRows(blockNum uint64, acctAddress string, err error, db *sql.DB, page int, pageSize int) (*sql.Rows, *sql.Rows, error) {
	var rows *sql.Rows
	var countRows *sql.Rows
	if blockNum != uint64(0) {
		if acctAddress != "" {
			if len(acctAddress) > 64 {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE block_num = ? AND sub_address = ? LIMIT ?,?", blockNum, acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE block_num = ? AND sub_address = ?", blockNum, acctAddress)
			} else if strings.HasPrefix(acctAddress, "CON") {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE block_num = ? AND contract_id = ? LIMIT ?,?", blockNum, acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE block_num = ? AND contract_id = ?", blockNum, acctAddress)
			} else {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE block_num = ? AND (from_addr = ? OR to_addr = ? )  LIMIT ?,?", blockNum, acctAddress, acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE block_num = ? AND (from_addr = ? OR to_addr = ?)", blockNum, acctAddress, acctAddress)
			}
		}

	} else {
		if acctAddress != "" {
			if len(acctAddress) > 64 {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE sub_address = ? LIMIT ?,?", acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE sub_address = ?", acctAddress)
			} else if strings.HasPrefix(acctAddress, "CON") {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE contract_id = ? LIMIT ?,?", acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE contract_id = ?", acctAddress)
			} else {
				rows, err = db.Query("SELECT * FROM tb_act_transaction WHERE (from_addr = ? OR to_addr = ? )  LIMIT ?,?", acctAddress, acctAddress, (page-1)*pageSize, pageSize)
				countRows, err = db.Query("SELECT count(*) FROM tb_act_transaction WHERE (from_addr = ? OR to_addr = ?)", acctAddress, acctAddress)
			}
		}
	}
	return rows, countRows, err
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
