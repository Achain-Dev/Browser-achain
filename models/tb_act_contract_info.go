package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbActContractInfo struct {
	Id         int    `orm:"column(id);auto" description:"唯一id"`
	ContractId string `orm:"column(contract_id);size(64)" description:"合约ID
"`
	Name         string    `orm:"column(name);size(64);null" description:"合约名称"`
	CoinType     string    `orm:"column(coin_type);size(70);null"`
	Bytecode     string    `orm:"column(bytecode)" description:"字节码"`
	Hash         string    `orm:"column(hash);size(64)" description:"字节码hash"`
	Owner        string    `orm:"column(owner);size(255)" description:"合约拥有者公钥"`
	OwnerAddress string    `orm:"column(owner_address);size(255)" description:"合约拥有者地址"`
	OwnerName    string    `orm:"column(owner_name);size(255);null" description:"合约拥有者名称"`
	Type         int8      `orm:"column(type);null" description:"0是其他 1是资产"`
	Description  string    `orm:"column(description);size(256);null" description:"合约描述"`
	RegTime      time.Time `orm:"column(reg_time);type(datetime)" description:"注册时间"`
	RegTrxId     string    `orm:"column(reg_trx_id);size(64)" description:"注册合约的交易id"`
	Balance      uint64    `orm:"column(balance)"`
	Circulation  int64     `orm:"column(circulation);null" description:"合约币发行总量"`
	Status       uint      `orm:"column(status)" description:"注册状态
0 - 销毁
1 - 临时
2 - 永久"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *TbActContractInfo) TableName() string {
	return "tb_act_contract_info"
}

func init() {
	orm.RegisterModel(new(TbActContractInfo))
}

// AddTbActContractInfo insert a new TbActContractInfo into database and returns
// last inserted Id on success.
func AddTbActContractInfo(m *TbActContractInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbActContractInfoById retrieves TbActContractInfo by Id. Returns error if
// Id doesn't exist
func GetTbActContractInfoById(id int) (v *TbActContractInfo, err error) {
	o := orm.NewOrm()
	v = &TbActContractInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbActContractInfo retrieves all TbActContractInfo matches certain condition. Returns empty list if
// no records exist
func GetAllTbActContractInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbActContractInfo))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TbActContractInfo
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTbActContractInfo updates TbActContractInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbActContractInfoById(m *TbActContractInfo) (err error) {
	o := orm.NewOrm()
	v := TbActContractInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbActContractInfo deletes TbActContractInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbActContractInfo(id int) (err error) {
	o := orm.NewOrm()
	v := TbActContractInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbActContractInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
