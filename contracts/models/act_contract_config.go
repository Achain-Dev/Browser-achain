package models

import (
	"Browser-achain/common"
	"log"
	"database/sql"
)

type TbActContractConfig struct {
	Id           int    `orm:"column(id);auto"`
	ContractId   string `orm:"column(contract_id);size(100);not null" description:"contract id"`
	ContractName *string `orm:"column(contract_name);size(255)" description:"contract name"`
	UrlIndex     int    `orm:"column(url_index)" description:"the meaning of the url"`
	Url          string `orm:"column(url);size(255);not null" description:"address of the url"`
	UrlName      *string `orm:"column(url_name);size(32);not null" description:"the name of url"`
	CreateTime   string `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime   string `orm:"column(update_time);type(timestamp)"`
}

// list all urls by contract id
func ListUrlsByContractId(contractId string) ([]TbActContractConfig, error) {
	db, err := common.GetDbConnection()

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

	return mappingDataToContractConfigList(rows)
}

func mappingDataToContractConfigList(rows *sql.Rows) ([]TbActContractConfig, error)  {
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

