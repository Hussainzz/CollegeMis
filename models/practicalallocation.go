package models

import (
	
	"github.com/astaxie/beego/orm"
	"time"
	
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