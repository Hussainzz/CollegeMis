package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
	"strings"
	"errors"
	"reflect"
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

func GetAllSubjectChoices(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(StudentSubjectChoices))
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

	var l []StudentSubjectChoices
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


