package models

import (
	
	"github.com/astaxie/beego/orm"
	"time"
	
)


type ProjectAllocation struct {
	Paid int `orm:"pk;auto" json:"Paid"`
	AcademicYear int `json:"AcademicYear"`
	ProjectCode string  `json:"ProjectCode"`
	ProjectName string `json:"projectName"`
	RollNo string `orm:"size(500)" json:"RollNo"`
	StaffMember *StaffMember `orm:"rel(fk)" json:"StaffMember"`
	SessionNumbers string `orm:"size(500)" json:"SessionNumbers"`
	AllocatedBy int `orm:"null" json:"AllocatedBy"`
	AllocatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"AllocatedAt"`
	Status int `orm:"null" json:"Status"`
}

func init(){
	orm.RegisterModel(new(ProjectAllocation))
}

func Projectallocation(pDetails *ProjectAllocation) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	attendance := &ProjectAllocation{}
	
	attendance.AcademicYear = pDetails.AcademicYear
	attendance.ProjectCode = pDetails.ProjectCode
	attendance.ProjectName = pDetails.ProjectName
	attendance.RollNo = pDetails.RollNo
	attendance.StaffMember = pDetails.StaffMember
	attendance.SessionNumbers = pDetails.SessionNumbers
	attendance.AllocatedBy = pDetails.AllocatedBy
	attendance.Status = 1


	o.Insert(attendance)

	success = "New Allocation Successfully"


	return success, nil
}

func (projectallocation ProjectAllocation) GetAllProjAllocation() (ret []ProjectAllocation) {
	o := orm.NewOrm()
	o.QueryTable("project_allocation").RelatedSel().OrderBy("-paid").All(&ret)
	return
}