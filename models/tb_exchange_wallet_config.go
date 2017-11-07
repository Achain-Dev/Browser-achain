package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbExchangeWalletConfig struct {
	Id             int       `orm:"column(id);auto"`
	ExchangeName   string    `orm:"column(exchange_name);size(64)" description:"交易所名称"`
	Status         int8      `orm:"column(status)" description:"0-有效,1-无效"`
	Url            string    `orm:"column(url);size(255)" description:"rpc调用的url"`
	RpcUser        string    `orm:"column(rpc_user);size(255)" description:"prc调用的username和password"`
	WalletName     string    `orm:"column(wallet_name);size(64)" description:"钱包名称"`
	WalletPassword string    `orm:"column(wallet_password);size(64)" description:"钱包密码"`
	PublicKey      string    `orm:"column(public_key);size(1000);null" description:"公钥"`
	ContractId     string    `orm:"column(contract_id);size(70);null"`
	ContractName   string    `orm:"column(contract_name);size(30);null"`
	CoinType       string    `orm:"column(coin_type);size(20)"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now_add"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *TbExchangeWalletConfig) TableName() string {
	return "tb_exchange_wallet_config"
}

func init() {
	orm.RegisterModel(new(TbExchangeWalletConfig))
}

// AddTbExchangeWalletConfig insert a new TbExchangeWalletConfig into database and returns
// last inserted Id on success.
func AddTbExchangeWalletConfig(m *TbExchangeWalletConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbExchangeWalletConfigById retrieves TbExchangeWalletConfig by Id. Returns error if
// Id doesn't exist
func GetTbExchangeWalletConfigById(id int) (v *TbExchangeWalletConfig, err error) {
	o := orm.NewOrm()
	v = &TbExchangeWalletConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbExchangeWalletConfig retrieves all TbExchangeWalletConfig matches certain condition. Returns empty list if
// no records exist
func GetAllTbExchangeWalletConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbExchangeWalletConfig))
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

	var l []TbExchangeWalletConfig
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

// UpdateTbExchangeWalletConfig updates TbExchangeWalletConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbExchangeWalletConfigById(m *TbExchangeWalletConfig) (err error) {
	o := orm.NewOrm()
	v := TbExchangeWalletConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbExchangeWalletConfig deletes TbExchangeWalletConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbExchangeWalletConfig(id int) (err error) {
	o := orm.NewOrm()
	v := TbExchangeWalletConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbExchangeWalletConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
