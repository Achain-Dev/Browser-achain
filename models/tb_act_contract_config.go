package models

import (
	"Browser-achain/util"
	"log"
)

type TbActContractConfig struct {
	Id           int    `orm:"column(id);auto"`
	ContractId   string `orm:"column(contract_id);size(100);not null" description:"合约id"`
	ContractName string `orm:"column(contract_name);size(255)" description:"合约名称"`
	UrlIndex     int    `orm:"column(url_index)" description:"url代表的含义"`
	Url          string `orm:"column(url);size(255);not null" description:"url地址"`
	UrlName      string `orm:"column(url_name);size(32);not null" description:"url名称"`
	CreateTime   string `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime   string `orm:"column(update_time);type(timestamp)"`
}

// list all urls by contract id
func ListUrlsByContractId(contractId string) ([]TbActContractConfig, error) {
	db, err := util.GetDbConnection()

	if err != nil {
		log.Fatal("ListUrlsByContractId|ERROR:", err)
		panic(err.Error())
	}

	defer db.Close()

	stat, err := db.Prepare("SELECT * FROM tb_act_contract_config WHERE contract_id = ?")

	if err != nil {
		log.Fatal("ListUrlsByContractId|ERROR:", err)
		panic(err.Error())
	}

	rows, err := stat.Query(contractId)

	if err != nil {
		log.Fatal("ListUrlsByContractId|ERROR:", err)
		panic(err.Error())
	}

	tbActContractConfigList := make([]TbActContractConfig, 0)

	for rows.Next() {
		var tbActContractConfig TbActContractConfig

		err := rows.Scan(
			&tbActContractConfig.Id,
			&tbActContractConfig.ContractId,
			&tbActContractConfig.ContractName,
			&tbActContractConfig.UrlIndex,
			&tbActContractConfig.Url,
			&tbActContractConfig.UrlName,
			&tbActContractConfig.CreateTime,
			&tbActContractConfig.UpdateTime,
		)

		if err != nil {
			log.Fatal("ListUrlsByContractId|ERROR:", err)
			panic(err)
			return tbActContractConfigList, err
		}
		tbActContractConfigList = append(tbActContractConfigList, tbActContractConfig)

	}

	return tbActContractConfigList, nil
}
