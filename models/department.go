package models

import (
	
	"github.com/astaxie/beego/orm"
	
	
)

type Department struct {
	Depid int `orm:"pk;auto" json:"Depid"`
	DepartmentName string `json:"DepartmentName"`
}

func init(){
	orm.RegisterModel(new(Department))
}

func AddingDepartment(dDetails *Department) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	deparment := &Department{}
	
	deparment.DepartmentName = dDetails.DepartmentName
	

	o.Insert(deparment)

	success = "New Department Added Successfully"


	return success, nil
}

func (department Department) GetAllDepartment() (ret []Department) {
	o := orm.NewOrm()
	o.QueryTable("department").All(&ret)
	return
}


