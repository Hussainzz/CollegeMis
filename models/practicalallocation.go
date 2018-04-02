package models

import (
	
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
	"errors"
	"reflect"
	
)

type PracticalAllocation struct {
	Prid int `orm:"pk;auto" json:"Prid"`
	AcademicYear int `json:"AcademicYear"`
	RollNo string `orm:"size(500)" json:"RollNo"`
	Semester string  `json:"Semester"`
	Batch string `orm:"null" json:"Batch"`
	StaffMember *StaffMember `orm:"rel(fk)" json:"StaffMember"`
	Division *Division `orm:"rel(fk)" json:"Division"`
	Subjects *Subjects `orm:"rel(fk)" json:"Subjects"`
	SessionNumbers string `orm:"size(500)" json:"SessionNumbers"`
	AllocatedBy int `orm:"null" json:"AllocatedBy"`
	AllocatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"AllocatedAt"`
	Status int `orm:"null" json:"Status"`
}

func init(){
	orm.RegisterModel(new(PracticalAllocation))
}

func Practicalallocation(pDetails *PracticalAllocation) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	attendance := &PracticalAllocation{}
	
	attendance.AcademicYear = pDetails.AcademicYear
	attendance.RollNo = pDetails.RollNo
	attendance.Semester = pDetails.Semester
	attendance.Batch = pDetails.Batch
	attendance.StaffMember = pDetails.StaffMember
	attendance.Division = pDetails.Division
	attendance.Subjects = pDetails.Subjects
	attendance.SessionNumbers = pDetails.SessionNumbers
	attendance.AllocatedBy = pDetails.AllocatedBy
	attendance.Status = 1


	o.Insert(attendance)

	success = "New Practical Allocation Successfully"


	return success, nil
}

func (practicalallocation PracticalAllocation) GetAllPractAllocation() (ret []PracticalAllocation) {
	o := orm.NewOrm()
	
	o.QueryTable("practical_allocation").RelatedSel().OrderBy("-prid").All(&ret)
	return
}

func GetAllPracticalAllocation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PracticalAllocation))
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

	var l []PracticalAllocation
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
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
