package models

import (
	"database/sql"
	"Browser-achain/common"
	"log"
)

type TbActBlock struct {
	Id             int       `orm:"column(id);auto"`
	BlockId        string    `orm:"column(block_id);size(64)" description:"block hash"`
	BlockNum       uint64    `orm:"column(block_num)" description:"block number"`
	BlockSize      uint64    `orm:"column(block_size)" description:"block size"`
	Previous       string    `orm:"column(previous);size(64)" description:"prev block id"`
	TrxDigest     *string    `orm:"column(trx_digest);size(128)" description:"Summary of transactions in blocks"`
	PrevSecret     string     `orm:"column(prev_secret);size(64)" description:"last secret"`
	NextSecretHash string    `orm:"column(next_secret_hash);size(64)" description:"current secret的hash"`
	RandomSeed     string    `orm:"column(random_seed);size(64)" description:"random send"`
	Signee         string    `orm:"column(signee);size(64)" description:"producer"`
	BlockTime      string `orm:"column(block_time);type(datetime)" description:"Block production time"`
	TransNum       uint      `orm:"column(trans_num)" description:"Number of transactions within the block"`
	TransAmount    uint64    `orm:"column(trans_amount)" description:"Total amount of transaction in the block"`
	TransFee       uint64    `orm:"column(trans_fee)" description:"Total transaction fee in block"`
	Status         int8      `orm:"column(status)"`
	CreateTime     string `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime     string `orm:"column(update_time);type(timestamp)"`
}

type ActBlockVO struct {
	Id int
	BlockId string
	BlockNum uint64
	BlockTime string
	TransNum uint
	TransAmount uint64
	Signee string
	BlockSize uint64
}

type ActBlockPageVO struct {
	ActBlockVOList []ActBlockVO
	CurrentPage uint
	PageSize uint
	TotalPage uint
	TotalRecords uint
}

// query block info
func BlockQueryByPage(signee string,page,pageSize int) (ActBlockPageVO,error) {
	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("BlockQueryByPage|ERROR:", err)
		panic(err.Error())
	}

	var rows *sql.Rows
	var countRows *sql.Rows

	if signee != "" {
		rows,err = db.Query("SELECT * FROM tb_act_block WHERE signee = ? LIMIT ?,?",signee,(page - 1) * pageSize,pageSize)
		countRows,err = db.Query("SELECT count(*) FROM tb_act_block WHERE signee = ?",signee)
	}else {
		rows,err = db.Query("SELECT * FROM tb_act_block LIMIT ?,?",(page - 1) * pageSize,pageSize)
		countRows,err = db.Query("SELECT count(*) FROM tb_act_block")
	}

	if err != nil {
		log.Fatal("BlockQueryByPage|ERROR:", err)
		panic(err.Error())
	}

	totalRecords := countNumber(countRows)
	var actBlockPageVO ActBlockPageVO
	if totalRecords == 0 {
		return actBlockPageVO,nil
	}

	totalPage := totalRecords / pageSize + 1
	if totalRecords % pageSize == 0 {
		totalPage = totalRecords / pageSize
	}
	actBlockPageVO.TotalPage = uint(totalPage)
	actBlockPageVO.CurrentPage = uint(page)
	actBlockPageVO.PageSize = uint(pageSize)
	actBlockPageVO.TotalRecords = uint(totalRecords)
	tbActBlockList, _ := mappingDataToBlockList(rows)
	if len(tbActBlockList) == 0 {
		return actBlockPageVO,nil
	}
	actBlockVOList := make([]ActBlockVO, len(tbActBlockList))

	for _,tbActBlock := range tbActBlockList {
		var actBlockVO ActBlockVO
		actBlockVO.Id = tbActBlock.Id
		actBlockVO.BlockId = tbActBlock.BlockId
		actBlockVO.BlockNum = tbActBlock.BlockNum
		actBlockVO.BlockTime = tbActBlock.BlockTime
		actBlockVO.TransNum = tbActBlock.TransNum
		actBlockVO.TransAmount = tbActBlock.TransAmount
		actBlockVO.Signee = tbActBlock.Signee
		actBlockVO.BlockSize = tbActBlock.BlockSize
		actBlockVOList = append(actBlockVOList,actBlockVO)
	}

	actBlockPageVO.ActBlockVOList = actBlockVOList
	return actBlockPageVO,nil
}

func BlockQueryByBlockId(blockId string) (*TbActBlock,error) {
	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("BlockQueryByBlockId|ERROR:", err)
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM tb_act_block WHERE block_id = ? LIMIT 1", blockId)
	if err != nil {
		log.Fatal("BlockQueryByBlockId|ERROR:", err)
		return nil,err
	}
	tbActBlockList, _ := mappingDataToBlockList(rows)
	if len(tbActBlockList) > 0 {
		return &tbActBlockList[0],nil
	}
	return nil,nil
}

func BlockQueryByBlockNum(blockNum int64) (*TbActBlock,error) {
	db, err := common.GetDbConnection()

	defer db.Close()

	if err != nil {
		log.Fatal("BlockQueryByBlockId|ERROR:", err)
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM tb_act_block WHERE block_num = ? LIMIT 1", blockNum)
	if err != nil {
		log.Fatal("BlockQueryByBlockId|ERROR:", err)
		return nil,err
	}
	tbActBlockList, _ := mappingDataToBlockList(rows)
	if len(tbActBlockList) > 0 {
		return &tbActBlockList[0],nil
	}
	return nil,nil
}




func mappingDataToBlockList(rows *sql.Rows) ([]TbActBlock,error)  {
	tbActBlockList := make([]TbActBlock, 0)
	for rows.Next() {
		var tbActBlock TbActBlock
		err := rows.Scan(
			&tbActBlock.Id,
			&tbActBlock.BlockId,
			&tbActBlock.BlockNum,
			&tbActBlock.BlockSize,
			&tbActBlock.Previous,
			&tbActBlock.TrxDigest,
			&tbActBlock.PrevSecret,
			&tbActBlock.NextSecretHash,
			&tbActBlock.RandomSeed,
			&tbActBlock.Signee,
			&tbActBlock.BlockTime,
			&tbActBlock.TransNum,
			&tbActBlock.TransAmount,
			&tbActBlock.TransFee,
			&tbActBlock.Status,
			&tbActBlock.CreateTime,
			&tbActBlock.UpdateTime,
		)
		if err != nil {
			log.Fatal("mappingDataToBlockList|ERROR:", err)
			panic(err)
			return tbActBlockList,err
		}
		tbActBlockList = append(tbActBlockList,tbActBlock)
	}

	return tbActBlockList,nil
}
