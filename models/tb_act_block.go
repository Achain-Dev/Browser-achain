package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbActBlock struct {
	Id             int       `orm:"column(id);auto"`
	BlockId        string    `orm:"column(block_id);size(64)" description:"区块hash"`
	BlockNum       uint64    `orm:"column(block_num)" description:"块号"`
	BlockSize      uint64    `orm:"column(block_size)" description:"块大小（字节）"`
	Previous       string    `orm:"column(previous);size(64)" description:"前一个块块id"`
	TrxDigest      string    `orm:"column(trx_digest);size(128)" description:"块中交易的摘要"`
	PrevSecret     string    `orm:"column(prev_secret);size(64)" description:"上轮secret"`
	NextSecretHash string    `orm:"column(next_secret_hash);size(64)" description:"本轮secret的hash"`
	RandomSeed     string    `orm:"column(random_seed);size(64)" description:"随机种子"`
	Signee         string    `orm:"column(signee);size(64)" description:"产块者"`
	BlockTime      time.Time `orm:"column(block_time);type(datetime)" description:"产块时间"`
	TransNum       uint      `orm:"column(trans_num)" description:"区块内交易数量"`
	TransAmount    uint64    `orm:"column(trans_amount)" description:"区块内交易总金额"`
	TransFee       uint64    `orm:"column(trans_fee)" description:"区块内交易总手续费"`
	Status         int8      `orm:"column(status)"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp)"`
}

func (t *TbActBlock) TableName() string {
	return "tb_act_block"
}

func init() {
	orm.RegisterModel(new(TbActBlock))
}

// AddTbActBlock insert a new TbActBlock into database and returns
// last inserted Id on success.
func AddTbActBlock(m *TbActBlock) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbActBlockById retrieves TbActBlock by Id. Returns error if
// Id doesn't exist
func GetTbActBlockById(id int) (v *TbActBlock, err error) {
	o := orm.NewOrm()
	v = &TbActBlock{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbActBlock retrieves all TbActBlock matches certain condition. Returns empty list if
// no records exist
func GetAllTbActBlock(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbActBlock))
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

	var l []TbActBlock
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

// UpdateTbActBlock updates TbActBlock by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbActBlockById(m *TbActBlock) (err error) {
	o := orm.NewOrm()
	v := TbActBlock{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbActBlock deletes TbActBlock by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbActBlock(id int) (err error) {
	o := orm.NewOrm()
	v := TbActBlock{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbActBlock{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
