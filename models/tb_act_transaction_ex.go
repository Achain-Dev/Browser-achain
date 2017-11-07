package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbActTransactionEx struct {
	Id         int       `orm:"column(id);auto" description:"唯一id"`
	TrxId      string    `orm:"column(trx_id);size(64)" description:"结果交易id"`
	OrigTrxId  string    `orm:"column(orig_trx_id);size(64)" description:"原始交易id"`
	FromAcct   string    `orm:"column(from_acct);size(64)" description:"发起账户"`
	FromAddr   string    `orm:"column(from_addr);size(64)" description:"发起地址"`
	ToAcct     string    `orm:"column(to_acct);size(64)" description:"接收账户"`
	ToAddr     string    `orm:"column(to_addr);size(64)" description:"接收地址"`
	Amount     uint64    `orm:"column(amount)" description:"金额"`
	Fee        uint      `orm:"column(fee)" description:"手续费"`
	Memo       string    `orm:"column(memo);size(3000);null" description:"备注"`
	TrxTime    time.Time `orm:"column(trx_time);type(datetime);null"`
	TrxType    int8      `orm:"column(trx_type)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *TbActTransactionEx) TableName() string {
	return "tb_act_transaction_ex"
}

func init() {
	orm.RegisterModel(new(TbActTransactionEx))
}

// AddTbActTransactionEx insert a new TbActTransactionEx into database and returns
// last inserted Id on success.
func AddTbActTransactionEx(m *TbActTransactionEx) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbActTransactionExById retrieves TbActTransactionEx by Id. Returns error if
// Id doesn't exist
func GetTbActTransactionExById(id int) (v *TbActTransactionEx, err error) {
	o := orm.NewOrm()
	v = &TbActTransactionEx{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbActTransactionEx retrieves all TbActTransactionEx matches certain condition. Returns empty list if
// no records exist
func GetAllTbActTransactionEx(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbActTransactionEx))
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

	var l []TbActTransactionEx
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

// UpdateTbActTransactionEx updates TbActTransactionEx by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbActTransactionExById(m *TbActTransactionEx) (err error) {
	o := orm.NewOrm()
	v := TbActTransactionEx{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbActTransactionEx deletes TbActTransactionEx by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbActTransactionEx(id int) (err error) {
	o := orm.NewOrm()
	v := TbActTransactionEx{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbActTransactionEx{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
