package models

import (
	"Browser-achain/common"
	"log"
	"database/sql"
)

type TbUserAddress struct {
	Id          int       `orm:"column(id);auto"`
	AccountName *string    `orm:"column(account_name);size(70);null"`
	UserAddress *string    `orm:"column(user_address);size(70)"`
	Balance     *int64     `orm:"column(balance)"`
	CoinType    string    `orm:"column(coin_type);size(70);null"`
	ContractId  *string    `orm:"column(contract_id);size(70);null"`
	TransNum    int       `orm:"column(trans_num)" description:"transaction number"`
	LastTrxTime string `orm:"column(last_trx_time);type(timestamp);null"`
	CreateTime  string `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime  string `orm:"column(update_time);type(timestamp)"`
}

func ListByAddress(address string) ([]TbUserAddress,error)  {
	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("ListByAddress|ERROR:", err)
		panic(err.Error())
	}

	stat, err := db.Prepare("SELECT * FROM tb_user_address WHERE user_address = ?")

	if err != nil {
		log.Fatal("ListByAddress|ERROR:", err)
		panic(err.Error())
	}

	rows, err := stat.Query(address)

	if err != nil {
		log.Fatal("ListByAddress|ERROR:", err)
		panic(err.Error())
	}

	return mappingDataToUserAddressList(rows)
}

func mappingDataToUserAddressList(rows *sql.Rows) ([]TbUserAddress,error) {
	tbUserAddressList := make([]TbUserAddress, 0)
	for rows.Next() {
		var tbUserAddress  TbUserAddress
		err := rows.Scan(
			&tbUserAddress.Id,
			&tbUserAddress.AccountName,
			&tbUserAddress.UserAddress,
			&tbUserAddress.Balance,
			&tbUserAddress.CoinType,
			&tbUserAddress.ContractId,
			&tbUserAddress.TransNum,
			&tbUserAddress.LastTrxTime,
			&tbUserAddress.CreateTime,
			&tbUserAddress.UpdateTime,
		)
		if err != nil {
			log.Fatal("ListByAddress|ERROR:", err)
			panic(err)
			return tbUserAddressList, err
		}
		tbUserAddressList = append(tbUserAddressList, tbUserAddress)

	}
	return tbUserAddressList,nil
}
