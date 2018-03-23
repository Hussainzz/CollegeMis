package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
)


type Subjects struct {
	Subid int `orm:"pk;auto" json:"Subid" `
	SubjectCode string `json:"SubjectCode"`
	SubjectName string `json:"SubjectName"`
	Course *Course `orm:"rel(fk)" json:"Course"`
	Semester string `json:"Semester"`
	Compulsory string `json:"Compulsory"`
}

type StudentSubjectChoices struct{
	Scid int `orm:"pk;auto" json:"Scid"`
	StudentDetails *StudentDetails `orm:"rel(fk)" json:"StudentDetails"`
	Subjets *Subjects `orm:"rel(fk)" json:"Subjects"`
	AcademicYear int `json:"AcademicYear"`
	Semester string `json:"Semester"`
	Course *Course `orm:"rel(fk)" json:"Course"`
}


func init(){
	orm.RegisterModel(new(Subjects), new(StudentSubjectChoices))
}

func AddingSubjects(sub *Subjects) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	subjects := &Subjects{}

	subjects.SubjectCode = sub.SubjectCode
	subjects.SubjectName = sub.SubjectName
	subjects.Course = sub.Course
	subjects.Semester = sub.Semester
	subjects.Compulsory = sub.Compulsory
	o.Insert(subjects)
	success = "New Subject Added Successfully"
	return success, nil
}


func (subjects Subjects) GetAllSubjects() (ret []Subjects) {
	o := orm.NewOrm()
	o.QueryTable("subjects").RelatedSel().OrderBy("-subid").All(&ret)
	return
}


func AddingStudentSubjectChoice(sub *StudentSubjectChoices) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	subjectchoice := &StudentSubjectChoices{}
	subjectchoice.StudentDetails = sub.StudentDetails
	subjectchoice.Subjets = sub.Subjets
	subjectchoice.AcademicYear = sub.AcademicYear
	subjectchoice.Semester = sub.Semester
	subjectchoice.Course = sub.Course

	o.Insert(subjectchoice)
	
	success = "Student Subject Choice Added Successfully"
	return success, nil
}