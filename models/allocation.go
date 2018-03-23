package models

import (
	
	"github.com/astaxie/beego/orm"
	"time"
	
)

type Allocation struct {
	Aaid int `orm:"pk;auto" json:"Aaid"`
	AcademicYear int `json:"AcademicYear"`
	Semester string `json:"Semester"`
	Division *Division `orm:"rel(fk)" json:"Division"`
	Subjects *Subjects `orm:"rel(fk)" json:"Subjects"`
	StaffMember *StaffMember `orm:"rel(fk)" json:"StaffMember"`
	SessionNumbers string `orm:"size(500)" json:"SessionNumbers"`
	AllocatedBy int `orm:"null" json:"AllocatedBy"`
	AllocatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"AllocatedAt"`
	Status int `orm:"null" json:"Status"`
}

func init(){
	orm.RegisterModel(new(Allocation))
}

func AddAllocation(aDetails *Allocation) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	attendance := &Allocation{}
	
	attendance.AcademicYear = aDetails.AcademicYear
	attendance.Semester = aDetails.Semester	
	attendance.Division = aDetails.Division
	attendance.Subjects = aDetails.Subjects
	attendance.StaffMember = aDetails.StaffMember
	attendance.SessionNumbers = aDetails.SessionNumbers
	attendance.AllocatedBy = aDetails.AllocatedBy
	attendance.Status = 1


	o.Insert(attendance)

	success = "New Allocation Successfully"


	return success, nil
}



func (allocation Allocation) GetAllAllocation() (ret []Allocation) {
	o := orm.NewOrm()
	o.QueryTable("allocation").RelatedSel().OrderBy("-aaid").All(&ret)
	return
}