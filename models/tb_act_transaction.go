package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbActTransaction struct {
	Id            int    `orm:"column(id);auto"`
	TrxId         string `orm:"column(trx_id);size(64)" description:"交易id"`
	BlockId       string `orm:"column(block_id);size(64)" description:"区块hash"`
	BlockNum      int64  `orm:"column(block_num)" description:"块号"`
	BlockPosition int    `orm:"column(block_position)" description:"交易在块中的位置"`
	TrxType       int    `orm:"column(trx_type)" description:"0 - 普通转账
1 - 代理领工资
2 - 注册账户
3 - 注册代理
10 - 注册合约
11 - 合约充值
12 - 合约升级
13 - 合约销毁
14 - 调用合约
15 - 合约出账
"`
	CoinType   string `orm:"column(coin_type);size(70);null"`
	ContractId string `orm:"column(contract_id);size(70);null"`
	FromAcct   string `orm:"column(from_acct);size(64)" description:"发起账号"`
	FromAddr   string `orm:"column(from_addr);size(64)" description:"发起地址"`
	ToAcct     string `orm:"column(to_acct);size(64)" description:"接收账号"`
	ToAddr     string `orm:"column(to_addr);size(64)" description:"接收地址"`
	SubAddress string `orm:"column(sub_address);size(70);null"`
	Amount     uint64 `orm:"column(amount)" description:"金额"`
	Fee        uint   `orm:"column(fee)" description:"手续费
如果是合约交易，包含gas消耗，注册保证金等"`
	Memo       string    `orm:"column(memo);null" description:"备注"`
	TrxTime    time.Time `orm:"column(trx_time);type(datetime)" description:"交易时间"`
	CalledAbi  string    `orm:"column(called_abi);size(6000);null" description:"调用的合约函数，非合约交易该字段为空"`
	AbiParams  string    `orm:"column(abi_params);size(6000);null" description:"调用合约函数时传入的参数，非合约交易该字段为空"`
	EventType  string    `orm:"column(event_type);size(32);null"`
	EventParam string    `orm:"column(event_param);size(1024);null"`
	ExtraTrxId string    `orm:"column(extra_trx_id);size(64);null" description:"结果交易id
仅针对合约交易"`
	IsCompleted uint8 `orm:"column(is_completed);null" description:"合约调用结果
0 - 成功
1- 失败"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *TbActTransaction) TableName() string {
	return "tb_act_transaction"
}

func init() {
	orm.RegisterModel(new(TbActTransaction))
}

// AddTbActTransaction insert a new TbActTransaction into database and returns
// last inserted Id on success.
func AddTbActTransaction(m *TbActTransaction) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbActTransactionById retrieves TbActTransaction by Id. Returns error if
// Id doesn't exist
func GetTbActTransactionById(id int) (v *TbActTransaction, err error) {
	o := orm.NewOrm()
	v = &TbActTransaction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbActTransaction retrieves all TbActTransaction matches certain condition. Returns empty list if
// no records exist
func GetAllTbActTransaction(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbActTransaction))
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

	var l []TbActTransaction
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

// UpdateTbActTransaction updates TbActTransaction by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbActTransactionById(m *TbActTransaction) (err error) {
	o := orm.NewOrm()
	v := TbActTransaction{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbActTransaction deletes TbActTransaction by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbActTransaction(id int) (err error) {
	o := orm.NewOrm()
	v := TbActTransaction{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbActTransaction{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
