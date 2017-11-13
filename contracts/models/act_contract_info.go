package models

import (
	"Browser-achain/common"
	"log"
	"database/sql"

)

const (
	Destroy = iota
	Temp
	Forever
)

type TbActContractInfo struct {
	Id           int
	ContractId   *string // contract id
	Name         *string // contract name
	CoinType     *string // contract type
	Bytecode     string  // byte
	Hash         string  // the hash of byte
	Owner        *string // Contract owner public key
	OwnerAddress *string // Address of the contract owner
	OwnerName    *string // Contract owner name
	Type         int8    // 0-other 1-asset
	Description  *string // Contract description
	RegTime      string  // register time
	RegTrxId     string  // The transaction id of the registered contract
	Balance      uint64  // balance
	Circulation  int64   // Total amount of contract currency issued
	Status       uint    // register status 0-destroy 1-temp 2-forever
	CreateTime   string
	UpdateTime   string
}

type ContractInfoVO struct {
	ContractName *string
	ContractId   *string
	CoinType     *string
	Status       uint
	CoinAddress  *string
	Circulation  int64
	RegisterTime string
	Coin         int
}

type ContractInfoPageVO struct {
	ActContractInfoList []TbActContractInfo
	ContractInfoVOList  []ContractInfoVO

	CurrentPage  uint
	PageSize     uint
	TotalPage    uint
	TotalRecords uint
}

// Query contract info by keyword
func ListContractInfoByKey(keyword string, contractStatus, page, perPage, queryType int) (ContractInfoPageVO, error) {
	var contractInfoPageVO ContractInfoPageVO
	contractInfoPageVO.CurrentPage = uint(page)
	contractInfoPageVO.PageSize = uint(perPage)

	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("ListContractInfoByKey|ERROR:", err)
		panic(err.Error())
	}

	var rows *sql.Rows
    var countRows *sql.Rows

	if keyword == "" {
		stmt, _ := db.Prepare("SELECT * FROM tb_act_contract_info WHERE status = ? ORDER BY reg_time DESC LIMIT ?,?")
		rows, err = stmt.Query(contractStatus,page - 1,perPage)
		countRows, err = db.Query("SELECT count(*) AS count FROM tb_act_contract_info WHERE status = ?", contractStatus)

	}else if queryType == 0 {
		stmt, _ := db.Prepare("SELECT * FROM tb_act_contract_info WHERE status = ? AND contract_id = ? ORDER BY reg_time DESC LIMIT ?,?")
		rows, err = stmt.Query(contractStatus,keyword,page - 1,perPage)
		countRows, err = db.Query("SELECT count(*) AS count FROM tb_act_contract_info WHERE status = ? AND contract_id = ?", contractStatus,keyword)
	}else if queryType == 1 {
		keyword = "%" + keyword + "%"
		stmt, _ := db.Prepare("SELECT * FROM tb_act_contract_info WHERE status = ? AND name LIKE  ? ORDER BY reg_time DESC LIMIT ?,?")
		rows, err = stmt.Query(contractStatus,keyword,page - 1,perPage)
		countRows, err = db.Query("SELECT count(*) AS count FROM tb_act_contract_info WHERE status = ? AND name LIKE  ?", contractStatus,keyword)
	}else {
		contractInfoPageVO.TotalPage = uint(0)
		contractInfoPageVO.TotalRecords = uint(0)
		return contractInfoPageVO,nil
	}

	if err != nil {
		log.Fatal("ListContractInfoByKey|ERROR:", err)
		panic(err.Error())
	}
	tbActContractInfoList, err := mappingDataToContractInfoList(rows)
	totalRecord := countNumber(countRows)

	totalPage := totalRecord / perPage + 1
	if totalRecord % perPage == 0 {
		totalPage = totalRecord / perPage
	}

	contractInfoPageVO.ActContractInfoList = tbActContractInfoList
	contractInfoPageVO.TotalPage = uint(totalPage)
	contractInfoPageVO.TotalRecords = uint(totalRecord)
	return contractInfoPageVO,nil
}

func countNumber(countRows *sql.Rows) int {
	count := 0
	for countRows.Next() {
       countRows.Scan(
       	&count,
	   )
	}
	return count
}

func mappingDataToContractInfoList(rows *sql.Rows) ([]TbActContractInfo,error) {

	tbActContractInfoList := make([]TbActContractInfo, 0)

	for rows.Next() {
		var tbActContractInfo TbActContractInfo

		err := rows.Scan(
			&tbActContractInfo.Id,
			&tbActContractInfo.ContractId,
			&tbActContractInfo.Name,
			&tbActContractInfo.CoinType,
			&tbActContractInfo.Bytecode,
			&tbActContractInfo.Hash,
			&tbActContractInfo.Owner,
			&tbActContractInfo.OwnerAddress,
			&tbActContractInfo.OwnerName,
			&tbActContractInfo.Type,
			&tbActContractInfo.Description,
			&tbActContractInfo.RegTime,
			&tbActContractInfo.RegTrxId,
			&tbActContractInfo.Balance,
			&tbActContractInfo.Circulation,
			&tbActContractInfo.Status,
			&tbActContractInfo.CreateTime,
			&tbActContractInfo.UpdateTime,
		)

		if err != nil {
			log.Fatal("mappingDataToContractInfoList|ERROR:", err)
			panic(err)
			return tbActContractInfoList, err
		}
		tbActContractInfoList = append(tbActContractInfoList, tbActContractInfo)
	}

	return tbActContractInfoList,nil
}
