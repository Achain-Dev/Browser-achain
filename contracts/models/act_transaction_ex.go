package models

import (
	"database/sql"
	"Browser-achain/common"
	"log"
)

type TbActTransactionEx struct {
	Id         int       `orm:"column(id);auto"`
	TrxId      *string    `orm:"column(trx_id);size(64)"`
	OrigTrxId  *string    `orm:"column(orig_trx_id);size(64)"`
	FromAcct   *string    `orm:"column(from_acct);size(64)"`
	FromAddr   *string    `orm:"column(from_addr);size(64)"`
	ToAcct     *string    `orm:"column(to_acct);size(64)"`
	ToAddr     *string    `orm:"column(to_addr);size(64)"`
	Amount     uint64    `orm:"column(amount)"`
	Fee        uint      `orm:"column(fee)"`
	Memo       *string    `orm:"column(memo);size(3000);null"`
	TrxTime    string  `orm:"column(trx_time);type(datetime);null"`
	TrxType    int8      `orm:"column(trx_type)"`
	CreateTime string  `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime string  `orm:"column(update_time);type(timestamp)"`
}

type ActTransactionExPage struct {
	ActTransactionExList []TbActTransactionEx
	TotalRecords int
	TotalPage int
	CurrentPage int
	PageSize int
}

func TransactionExQuery(originTrxId string,page,pageSize int)  (ActTransactionExPage,error) {

	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("TransactionExQuery|ERROR:", err)
		panic(err.Error())
	}

	var rows *sql.Rows
	var countRows *sql.Rows

	if originTrxId != "" {
		rows,err = db.Query("SELECT * FROM tb_act_transaction_ex WHERE orig_trx_id = ? LIMIT ?,?",originTrxId,(page - 1) * pageSize,pageSize)
		countRows,err = db.Query("SELECT count(*) FROM tb_act_transaction_ex WHERE orig_trx_id = ?",originTrxId)
	}else {
		rows,err = db.Query("SELECT * FROM tb_act_transaction_ex LIMIT ?,?",(page - 1) * pageSize,pageSize)
		countRows,err = db.Query("SELECT count(*) FROM tb_act_transaction_ex")
	}

	if err != nil {
		log.Fatal("TransactionExQuery|ERROR:", err)
		panic(err.Error())
	}

	totalRecords := countNumber(countRows)
	var actTransactionExPage ActTransactionExPage
	if totalRecords == 0 {
		return actTransactionExPage,nil
	}

	totalPage := totalRecords / pageSize + 1
	if totalRecords % pageSize == 0 {
		totalPage = totalRecords / pageSize
	}
	actTransactionExPage.TotalPage = totalPage
	actTransactionExPage.CurrentPage = page
	actTransactionExPage.PageSize = pageSize
	actTransactionExPage.TotalRecords = totalRecords
	tbActContractInfoList, _ := mappingDataToTransactionExList(rows)
	actTransactionExPage.ActTransactionExList = tbActContractInfoList
	return actTransactionExPage,nil
}

func mappingDataToTransactionExList(rows *sql.Rows) ([]TbActTransactionEx,error)  {

	tbActContractInfoList := make([]TbActTransactionEx, 0)
	for rows.Next() {
		var tbActTransactionEx TbActTransactionEx

		err := rows.Scan(
			&tbActTransactionEx.Id,
				&tbActTransactionEx.TrxId,
					&tbActTransactionEx.OrigTrxId,
						&tbActTransactionEx.FromAcct,
						&tbActTransactionEx.FromAddr,
							&tbActTransactionEx.ToAcct,
								&tbActTransactionEx.ToAddr,
									&tbActTransactionEx.Amount,
										&tbActTransactionEx.Fee,
											&tbActTransactionEx.Memo,
												&tbActTransactionEx.TrxTime,
													&tbActTransactionEx.TrxType,
														&tbActTransactionEx.CreateTime,
															&tbActTransactionEx.UpdateTime,
		)
		if err != nil {
			log.Fatal("mappingDataToTransactionExList|ERROR:", err)
			panic(err)
			return tbActContractInfoList,err
		}
		tbActContractInfoList = append(tbActContractInfoList,tbActTransactionEx)
	}

	return tbActContractInfoList,nil
}








