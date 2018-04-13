package controllers

import (
	"fmt"
	"CollegeMis/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
	//"reflect"
	"strings"
	"errors"
)

// Operations about Student
type StudentController struct{
	beego.Controller
}
// @Title Student
// @Description Adding Student
// @Param	body		body 	models.StudentDetails	true		"body for user content"
// @Success 200 {int} models.StudentDetails.Sid
// @Failure 403 body is empty
// @router / [post]
func (r *StudentController) Post(){

	var sds models.StudentDetails
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &sds); err == nil {
		if success, err := models.AddStudentDetails(&sds); err == nil {
			r.Ctx.Output.SetStatus(201) 
			//errorType := "null"
			
			r.Data["json"] = success
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()
}



// Operations about Student
type StudentAllController struct{
	beego.Controller
}


// @Title GetAll
// @Description get all StudentDetails
// @Success 200 {StudentDetails} models.StudentDetails
// @router / [get]
func(sd *StudentAllController) Get(){
	studentdetails := &models.StudentDetails{}
	s := studentdetails.GetAllStudentDetails()
	sd.Data["json"] = s
	sd.ServeJSON()
}
// Operations about Student
type StudentViewController struct{
	beego.Controller
}
// @Title Get
// @Description find Student by Sid
// @Param	Sid		path 	string	true		"the Sid you want to get"
// @Success 200 {object} models.StudentDetails
// @Failure 403 :Sid is empty
// @router /:Sid [get]
func (s *StudentViewController) Get(){
	Sid,_ := strconv.Atoi(s.Ctx.Input.Param(":Sid"))
	var ret []models.StudentDetails
	o := orm.NewOrm()
	
	o.QueryTable("student_details").Filter("sid__exact", Sid).RelatedSel().All(&ret)

	s.Data["json"] = ret
	s.ServeJSON()
}

// Operations about Student
type StudentEditController struct{
	beego.Controller
}


// @Title Update
// @Description update the Student
// @Param	Sid		path 	string	true		"The Sid you want to update"
// @Param	body		body 	models.StudentDetails	true		"body for student content"
// @Success 200 {Sid} models.StudentDetails
// @Failure 403 :Sid is not int
// @router /:Sid [put]
func (this *StudentEditController) Put(){
	o := orm.NewOrm()
	Sid,_ := strconv.Atoi(this.Ctx.Input.Param(":Sid"))
	var sd models.StudentDetails

	json.Unmarshal(this.Ctx.Input.RequestBody, &sd)

	ss := models.StudentDetails{Sid: Sid}
		if o.Read(&ss) == nil {
	ss.RollNo = sd.RollNo
	ss.Course = sd.Course
	ss.AdmissionYear = sd.AdmissionYear
	ss.Division = sd.Division
	ss.Semester = sd.Semester

	ss.FirstName = sd.FirstName
	ss.MiddleName = sd.MiddleName
	ss.LastName = sd.LastName
	ss.Gender = sd.Gender
	ss.Dob = sd.Dob
	ss.PlaceOfBirth = sd.PlaceOfBirth
	ss.MaritalStatus = sd.MaritalStatus
	ss.Religon = sd.Religon
	ss.Nationality = sd.Nationality
	ss.Category = sd.Category
	ss.MobileNumber = sd.MobileNumber
	ss.BloodGroup = sd.BloodGroup
	ss.BankName = sd.BankName
	ss.AccountNumber = sd.AccountNumber
	ss.BranchAddress = sd.BranchAddress


	ss.LandLineNumber = sd.LandLineNumber
	ss.ParentMobileNumber = sd.ParentMobileNumber
	ss.Email = sd.Email
	ss.ResidentialAddress = sd.ResidentialAddress
	ss.FatherName = sd.FatherName
	ss.FatherStatus = sd.FatherStatus
	ss.FatherOccupation = sd.FatherOccupation
	ss.MotherName = sd.MotherName
	ss.MotherStatus = sd.MotherStatus
	ss.MotherOccupation = sd.MotherOccupation
	ss.ParentEmail = sd.ParentEmail
	ss.SiblingName = sd.SiblingName
	ss.GuardianName = sd.GuardianName
	ss.GuardianOccupation = sd.GuardianOccupation
	ss.GuardianRelation = sd.GuardianRelation
	ss.GuardianAddress = sd.GuardianAddress

	ss.SscBoard = sd.SscBoard
	ss.SscSeatNo = sd.SscSeatNo
	ss.SscMonthOfPassing = sd.SscMonthOfPassing
	ss.SscYearOfPassing = sd.SscYearOfPassing
	ss.SscTotalMarks = sd.SscTotalMarks
	ss.SscMarkOutOf = sd.SscMarkOutOf
	ss.SscPercentage = sd.SscPercentage
	ss.SscClassObtained = sd.SscClassObtained
	
	ss.HsscBoard = sd.HsscBoard
	ss.HsscSeatNo = sd.HsscSeatNo
	ss.HsscStream = sd.HsscStream
	ss.HsscMonthOfPassing = sd.HsscMonthOfPassing
	ss.HsscYearOfPassing = sd.HsscYearOfPassing
	ss.HsscTotalMarks = sd.HsscTotalMarks
	ss.HsscMarkOutOf = sd.HsscMarkOutOf
	ss.HsscPercentage = sd.HsscPercentage
	ss.HsscClassObtained = sd.HsscClassObtained

	ss.Scholorship = sd.Scholorship
	ss.CoCurricularActivities = sd.CoCurricularActivities
	ss.ExtraCurricularActivities = sd.ExtraCurricularActivities
	ss.SportsName = sd.SportsName
	ss.SportsLevelOfParticipation = sd.SportsLevelOfParticipation
			
   			 if num, err := o.Update(&ss); err == nil {
					fmt.Println(num)
        			this.Data["json"] = "Updated"
    			}
		}
	this.ServeJSON()	
}


// Operations about Student
type StudentDeleteController struct{
	beego.Controller
}


// @Title Delete
// @Description delete the Student
// @Param	Sid		path 	string	true		"The Sid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Sid is empty
// @router /:Sid [delete]
func (this *StudentDeleteController) Delete(){
	o := orm.NewOrm()
	Sid, _ := strconv.Atoi(this.Ctx.Input.Param(":Sid"))
			studentD := &models.StudentDetails{Sid:Sid}
			
			
			if num, err := o.Delete(studentD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type StudentSearchController struct{
	beego.Controller
}


// @Title Student
// @Description Searching Student
// @Param	body		body 	models.StudentDetails	true		"body for user content"
// @Success 200 {int} models.StudentDetails.Sid
// @Failure 403 body is empty
// @router / [post]
func (r *StudentSearchController) Post(){
	
	var Key models.Keyword


	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &Key); err == nil {
		if ret, success, err := models.SearchStudent(&Key); err == nil {
			r.Ctx.Output.SetStatus(201) 
		
			fmt.Println(success)
			
			if len(ret) <= 0{
				r.Data["json"] = "Nothing Found"
			}else{
				r.Data["json"] = ret
			}
			
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()


}




type StudentSController struct{
	beego.Controller
}
// GetAll ...
// @Title Get All
// @Description get StudentDetails
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.StudentDetails
// @Failure 403
// @router / [get]
func (c *StudentSController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllStudentDetails(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

type CourseController struct{
	beego.Controller
}
// @Title Course
// @Description Adding Course
// @Param	body		body 	models.Course	true		"body for user content"
// @Success 200 {int} models.Course.Cid
// @Failure 403 body is empty
// @router / [post]
func (r *CourseController) Post(){

	var cd models.Course
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddingCourse(&cd); err == nil {
			
			r.Data["json"] = success
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()
}

type CourseAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all Course
// @Success 200 {Course} models.Course
// @router / [get]
func(cc *CourseAllController) Get(){
	course := &models.Course{}
	c := course.GetAllCourse()
	cc.Data["json"] = c
	cc.ServeJSON()
}

type CourseDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Course
// @Param	Cid		path 	string	true		"The Cid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Cid is empty
// @router /:Cid [delete]
func (this *CourseDeleteController) Delete(){
	o := orm.NewOrm()
	Cid, _ := strconv.Atoi(this.Ctx.Input.Param(":Cid"))
			CourseD := &models.Course{Cid:Cid}
			
			
			if num, err := o.Delete(CourseD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type CourseEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the 
// @Param	Cid		path 	string	true		"The Cid you want to update"
// @Param	body		body 	models.Course	true		"body for Department content"
// @Success 200 {Cid} models.Course
// @Failure 403 :Cid is not int
// @router /:Cid [put]
func (this *CourseEditController) Put(){
	o := orm.NewOrm()
	Cid,_ := strconv.Atoi(this.Ctx.Input.Param(":Cid"))
	var nd models.Course

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.Course{Cid: Cid}
		if o.Read(&nn) == nil {
	
			nn.CourseShortName = nd.CourseShortName
			nn.CourseFullName = nd.CourseFullName
			nn.CourseType = nd.CourseType
			nn.RollPrefix = nd.RollPrefix
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
}


type DivisionController struct{
	beego.Controller
}

// @Title Division
// @Description Adding Divisions
// @Param	body		body 	models.Division	true		"body for user content"
// @Success 200 {int} models.Division.Did
// @Failure 403 body is empty
// @router / [post]
func (r *DivisionController) Post(){

	var dd models.Division
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &dd); err == nil {
		if success, err := models.AddingDivision(&dd); err == nil {
			r.Ctx.Output.SetStatus(201) 
			
			r.Data["json"] = success
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()
}



type DivisionAllController struct{
	beego.Controller
}

// @Title GetAll
// @Description get all Division
// @Success 200 {Division} models.Division
// @router / [get]
func(dd *DivisionAllController) Get(){
	division := &models.Division{}
	d := division.GetAllDivision()
	dd.Data["json"] = d
	dd.ServeJSON()
}



type DivisionViewController struct{
	beego.Controller
}
// @Title Get
// @Description find Division by Cid
// @Param	Cid		path 	string	true		"the Cid you want to get"
// @Success 200 {object} models.Division
// @Failure 403 :Cid is empty
// @router /:Cid [get]
func (s *DivisionViewController) Get(){
	Cid,_ := strconv.Atoi(s.Ctx.Input.Param(":Cid"))
	//student := &models.StudentDetails{}
	
	var ret []models.Division
	o := orm.NewOrm()
	
	o.QueryTable("division").Filter("course_id__exact", Cid).RelatedSel().All(&ret)

	s.Data["json"] = ret
	s.ServeJSON()

}

type DivisionDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Division
// @Param	Did		path 	string	true		"The Did you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Did is empty
// @router /:Did [delete]
func (this *DivisionDeleteController) Delete(){
	o := orm.NewOrm()
	Did, _ := strconv.Atoi(this.Ctx.Input.Param(":Did"))
			DivisionD := &models.Division{Did:Did}
			
			
			if num, err := o.Delete(DivisionD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type DivisionEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the 
// @Param	Did		path 	string	true		"The Did you want to update"
// @Param	body		body 	models.Division	true		"body for Department content"
// @Success 200 {Did} models.Division
// @Failure 403 :Did is not int
// @router /:Did [put]
func (this *DivisionEditController) Put(){
	o := orm.NewOrm()
	Did,_ := strconv.Atoi(this.Ctx.Input.Param(":Did"))
	var nd models.Division

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.Division{Did: Did}
		if o.Read(&nn) == nil {
	
			nn.DivisionName = nd.DivisionName
			nn.Course = nd.Course
			nn.YearClass = nd.YearClass
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
}