package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbActWithdraw struct {
	Id          int       `orm:"column(id);auto"`
	TrxId       string    `orm:"column(trx_id);size(64)" description:"交易id"`
	WalletName  string    `orm:"column(wallet_name);size(64)" description:"钱包名"`
	AssetSymbol string    `orm:"column(asset_symbol);size(64)" description:"资产类型"`
	FromAcct    string    `orm:"column(from_acct);size(64)" description:"发起账号"`
	FromAddr    string    `orm:"column(from_addr);size(70);null" description:"发起地址"`
	ToAcct      string    `orm:"column(to_acct);size(64);null" description:"接收账号"`
	ToAddr      string    `orm:"column(to_addr);size(70)" description:"接收地址"`
	Amount      string    `orm:"column(amount);size(100)" description:"金额"`
	Memo        string    `orm:"column(memo);size(500);null" description:"备注"`
	CreateTime  time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp)"`
	BlockTrxId  string    `orm:"column(block_trx_id);size(70);null"`
}

func (t *TbActWithdraw) TableName() string {
	return "tb_act_withdraw"
}

func init() {
	orm.RegisterModel(new(TbActWithdraw))
}

// AddTbActWithdraw insert a new TbActWithdraw into database and returns
// last inserted Id on success.
func AddTbActWithdraw(m *TbActWithdraw) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbActWithdrawById retrieves TbActWithdraw by Id. Returns error if
// Id doesn't exist
func GetTbActWithdrawById(id int) (v *TbActWithdraw, err error) {
	o := orm.NewOrm()
	v = &TbActWithdraw{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbActWithdraw retrieves all TbActWithdraw matches certain condition. Returns empty list if
// no records exist
func GetAllTbActWithdraw(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbActWithdraw))
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

	var l []TbActWithdraw
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

// UpdateTbActWithdraw updates TbActWithdraw by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbActWithdrawById(m *TbActWithdraw) (err error) {
	o := orm.NewOrm()
	v := TbActWithdraw{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbActWithdraw deletes TbActWithdraw by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbActWithdraw(id int) (err error) {
	o := orm.NewOrm()
	v := TbActWithdraw{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbActWithdraw{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
