package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type JOBSTATUSTRACELOG struct {
	Id             int       `orm:"column(id);pk"`
	JobName        string    `orm:"column(job_name);size(100)"`
	OriginalTaskId string    `orm:"column(original_task_id);size(255)"`
	TaskId         string    `orm:"column(task_id);size(255)"`
	SlaveId        string    `orm:"column(slave_id);size(50)"`
	Source         string    `orm:"column(source);size(50)"`
	ExecutionType  string    `orm:"column(execution_type);size(20)"`
	ShardingItem   string    `orm:"column(sharding_item);size(100)"`
	State          string    `orm:"column(state);size(20)"`
	Message        string    `orm:"column(message);size(4000);null"`
	CreationTime   time.Time `orm:"column(creation_time);type(timestamp);null"`
}

func (t *JOBSTATUSTRACELOG) TableName() string {
	return "JOB_STATUS_TRACE_LOG"
}

func init() {
	orm.RegisterModel(new(JOBSTATUSTRACELOG))
}

// AddJOBSTATUSTRACELOG insert a new JOBSTATUSTRACELOG into database and returns
// last inserted Id on success.
func AddJOBSTATUSTRACELOG(m *JOBSTATUSTRACELOG) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetJOBSTATUSTRACELOGById retrieves JOBSTATUSTRACELOG by Id. Returns error if
// Id doesn't exist
func GetJOBSTATUSTRACELOGById(id int) (v *JOBSTATUSTRACELOG, err error) {
	o := orm.NewOrm()
	v = &JOBSTATUSTRACELOG{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllJOBSTATUSTRACELOG retrieves all JOBSTATUSTRACELOG matches certain condition. Returns empty list if
// no records exist
func GetAllJOBSTATUSTRACELOG(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(JOBSTATUSTRACELOG))
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

	var l []JOBSTATUSTRACELOG
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

// UpdateJOBSTATUSTRACELOG updates JOBSTATUSTRACELOG by Id and returns error if
// the record to be updated doesn't exist
func UpdateJOBSTATUSTRACELOGById(m *JOBSTATUSTRACELOG) (err error) {
	o := orm.NewOrm()
	v := JOBSTATUSTRACELOG{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteJOBSTATUSTRACELOG deletes JOBSTATUSTRACELOG by Id and returns error if
// the record to be deleted doesn't exist
func DeleteJOBSTATUSTRACELOG(id int) (err error) {
	o := orm.NewOrm()
	v := JOBSTATUSTRACELOG{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&JOBSTATUSTRACELOG{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
