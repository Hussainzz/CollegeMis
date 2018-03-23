package models

import (
	
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strings"
	"strconv"
	"time"	
	"github.com/astaxie/beego"
	"reflect"
	"strings"
	"errors"
)

type StudentDetails struct{
	Sid int `orm:"pk;auto" json:"Sid"`
	RegistrationId string `json:"RegistrationId"`
	RollNo string `json:"RollNo"`
	Course *Course `orm:"rel(fk)" json:"Course"`
	AdmissionYear int `json:"AdmissionYear"`
	Division *Division `orm:"rel(fk)" json:"Division"`
	Semester string `json:"Semester"`

	FirstName string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName string `json:"LastName"`
	Gender string `json:"Gender"`
	Dob string `json:"Dob"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	MaritalStatus string `json:"MaritalStatus"`
	Religon string `json:"Religon"`
	Nationality string `json:"Nationality"`
	Category string `json:"Category"`
	MobileNumber string `json:"MobileNumber"`
	BloodGroup string `json:"BloodGroup"`
	BankName string `json:"BankName"`
	AccountNumber string `json:"AccountNumber"`
	BranchAddress string `json:"BranchAddress"`

	LandLineNumber string `json:"LandLineNumber"`
	ParentMobileNumber string `json:"ParentMobileNumber"`
	Email string `json:"Email"`
	ResidentialAddress string `orm:"size(1500)" json:"ResidentialAddress"`
	FatherName string `json:"FatherName"`
	FatherStatus string `json:"FatherStatus"`
	FatherOccupation string `json:"FatherOccupation"`
	MotherName string `json:"MotherName"`
	MotherStatus string `json:"MotherStatus"`
	MotherOccupation string `json:"MotherOccupation"`
	ParentEmail string `json:"ParentEmail"`
	SiblingName string `json:"SiblingName"`
	GuardianName string `json:"GuardianName"`
	GuardianOccupation string `json:"GuardianOccupation"`
	GuardianRelation string `json:"GuardianRelation"`
	GuardianAddress string `orm:"size(1500)" json:"GuardianAddress"`

	SscBoard string `json:"SscBoard"`
	SscSeatNo string `json:"SscSeatNo"`
	SscMonthOfPassing string `json:"SscMonthOfPassing"`
	SscYearOfPassing int ` json:"SscYearOfPassing"`
	SscTotalMarks int `json:"SscTotalMarks"`
	SscMarkOutOf int `json:"SscMarkOutOf"`
	SscPercentage float64 `json:"SscPercentage"`
	SscClassObtained string `json:"SscClassObtained"`

	HsscBoard string `json:"HsscBoard"`
	HsscStream string `json:"HsscStream"`
	HsscSeatNo string `json:"HsscSeatNo"`
	HsscMonthOfPassing string `json:"HsscMonthOfPassing"`
	HsscYearOfPassing int ` json:"HsscYearOfPassing"`
	HsscTotalMarks int `json:"HsscTotalMarks"`
	HsscMarkOutOf int `json:"HsscMarkOutOf"`
	HsscPercentage float64 `json:"HsscPercentage"`
	HsscClassObtained string `json:"HsscClassObtained"`

	Scholorship string `json:"Scholorship"`
	CoCurricularActivities string `orm:"size(600)" json:"CoCurricularActivities"`
	ExtraCurricularActivities string `orm:"size(600)" json:"ExtraCurricularActivities"`
	SportsName string `json:"SportsName"`
	SportsLevelOfParticipation string ` json:"SportsLevelOfParticipation"`
}

type Course struct {
	Cid int `orm:"pk;auto" json:"Cid"`
	CourseShortName string `json:"CourseShortName"`
 	CourseFullName string `json:"CourseFullName"`
	CourseType string `json:"CourseType"`
	RollPrefix string `json:"RollPrefix"`
}



type Division struct{
	Did int `orm:"pk;auto" json:"Did"`
	DivisionName string `json:"DivisionName"`
	Course *Course `orm:"rel(fk)" json:"Course"`
	YearClass string `json:"YearClass"`
}

func init(){
	orm.RegisterModel(new(StudentDetails), new(Course), new(Division))
}

func AddStudentDetails(sds *StudentDetails)(string, error){
	o := orm.NewOrm()
	o.Using("default")
	sd := new(StudentDetails)
	sd.RegistrationId = strconv.FormatInt(time.Now().Unix(), 10)
	sd.RollNo = sds.RollNo
	sd.Course = sds.Course
	sd.AdmissionYear = sds.AdmissionYear
	sd.Division = sds.Division
	sd.Semester = sds.Semester

	sd.FirstName = sds.FirstName
	sd.MiddleName = sds.MiddleName
	sd.LastName = sds.LastName
	sd.Gender = sds.Gender
	sd.Dob = sds.Dob
	sd.PlaceOfBirth = sds.PlaceOfBirth
	sd.MaritalStatus = sds.MaritalStatus
	sd.Religon = sds.Religon
	sd.Nationality = sds.Nationality
	sd.Category = sds.Category
	sd.MobileNumber = sds.MobileNumber
	sd.BloodGroup = sds.BloodGroup
	sd.BankName = sds.BankName
	sd.AccountNumber = sds.AccountNumber
	sd.BranchAddress = sds.BranchAddress


	sd.LandLineNumber = sds.LandLineNumber
	sd.ParentMobileNumber = sds.ParentMobileNumber
	sd.Email = sds.Email
	sd.ResidentialAddress = sds.ResidentialAddress
	sd.FatherName = sds.FatherName
	sd.FatherStatus = sds.FatherStatus
	sd.FatherOccupation = sds.FatherOccupation
	sd.MotherName = sds.MotherName
	sd.MotherStatus = sds.MotherStatus
	sd.MotherOccupation = sds.MotherOccupation
	sd.ParentEmail = sds.ParentEmail
	sd.SiblingName = sds.SiblingName
	sd.GuardianName = sds.GuardianName
	sd.GuardianOccupation = sds.GuardianOccupation
	sd.GuardianRelation = sds.GuardianRelation
	sd.GuardianAddress = sds.GuardianAddress

	sd.SscBoard = sds.SscBoard
	sd.SscSeatNo = sds.SscSeatNo
	sd.SscMonthOfPassing = sds.SscMonthOfPassing
	sd.SscYearOfPassing = sds.SscYearOfPassing
	sd.SscTotalMarks = sds.SscTotalMarks
	sd.SscMarkOutOf = sds.SscMarkOutOf
	sd.SscPercentage = sds.SscPercentage
	sd.SscClassObtained = sds.SscClassObtained
	
	sd.HsscBoard = sds.HsscBoard
	sd.HsscSeatNo = sds.HsscSeatNo
	sd.HsscStream = sds.HsscStream
	sd.HsscMonthOfPassing = sds.HsscMonthOfPassing
	sd.HsscYearOfPassing = sds.HsscYearOfPassing
	sd.HsscTotalMarks = sds.HsscTotalMarks
	sd.HsscMarkOutOf = sds.HsscMarkOutOf
	sd.HsscPercentage = sds.HsscPercentage
	sd.HsscClassObtained = sds.HsscClassObtained

	sd.Scholorship = sds.Scholorship
	sd.CoCurricularActivities = sds.CoCurricularActivities
	sd.ExtraCurricularActivities = sds.ExtraCurricularActivities
	sd.SportsName = sds.SportsName
	sd.SportsLevelOfParticipation = sds.SportsLevelOfParticipation
	o.Insert(sd)
	success := "ThankYou For Registring You Registration Id. : "+ sd.RegistrationId
return success, nil
	
}

func SearchStudent(key *Keyword) ( []StudentDetails, string, error){
	
	o := orm.NewOrm()
	
	query := o.QueryTable("student_details")

	var ret []StudentDetails

	query.Filter("first_name", key.Key).RelatedSel().All(&ret)
	query.Filter("course_id", key.Key).RelatedSel().All(&ret)
	query.Filter("AdmissionYear", key.Key).RelatedSel().All(&ret)
	query.Filter("registration_id",key.Key).RelatedSel().All(&ret)

		success := "Queryed Students"
		return  ret, success, nil
}

func (studentdetails StudentDetails) GetAllStudentDetails() (ret []StudentDetails) {
	o := orm.NewOrm()
	o.QueryTable("student_details").RelatedSel().OrderBy("-sid").All(&ret)
	return
}

func (sd *StudentDetails) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(sd); err != nil {
		beego.Info(err)
	}

	return
}


func AddingCourse(cDetails *Course) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	course := &Course{}
	
	course.CourseShortName = cDetails.CourseShortName
	course.CourseFullName = cDetails.CourseFullName
	course.CourseType = cDetails.CourseType
	course.RollPrefix = cDetails.RollPrefix
	o.Insert(course)

	success = "New Course Added Successfully"
	return success, nil
}


func AddingDivision(dDetails *Division) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	division := &Division{}
	
	division.DivisionName = dDetails.DivisionName
	division.Course = dDetails.Course
	division.YearClass = dDetails.YearClass

	o.Insert(division)

	success = "New Division Added Successfully"


	return success, nil
}

func (course Course) GetAllCourse() (ret []Course) {
	o := orm.NewOrm()
	o.QueryTable("course").All(&ret)
	return
}

func (division Division) GetAllDivision() (ret []Division) {
	o := orm.NewOrm()
	o.QueryTable("division").RelatedSel().All(&ret)
	return
}





func GetAllStudentDetails(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(StudentDetails))
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

	var l []StudentDetails
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