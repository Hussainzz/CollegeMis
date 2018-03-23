package controllers

import (
	"CollegeMis/models"
	"encoding/json"
	"errors"
	"strconv"
	"fmt"
	"strings"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)


type StaffMemberController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create StaffMember
// @Param	body		body 	models.StaffMember	true		"body for StaffMember content"
// @Success 201 {int} models.StaffMember
// @Failure 403 body is empty
// @router / [post]
func (r *StaffMemberController) Post() {
	var nd models.StaffMember
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &nd); err == nil {
		if success, err := models.AddingStaffMember(&nd); err == nil {
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


type StaffMemberAllController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Staff
// @Success 200 {object} models.StaffMember
// @router / [get]
func(sd *StaffMemberAllController) Get(){
	staffdetails := &models.StaffMember{}
	s := staffdetails.GetAllStaffDetails()
	sd.Data["json"] = s
	sd.ServeJSON()
}



type StaffMemberViewController struct {
	beego.Controller
}
// @Title Get
// @Description find Staff by StaffId
// @Param	StaffId		path 	string	true		"the StaffId you want to get"
// @Success 200 {object} models.StaffMember
// @Failure 403 :StaffId is empty
// @router /:StaffId [get]
func (s *StaffMemberViewController) Get(){
	StaffId,_ := strconv.Atoi(s.Ctx.Input.Param(":StaffId"))

	var ret []models.StaffMember
	o := orm.NewOrm()
	
	o.QueryTable("staff_member").Filter("staff_id__exact", StaffId).RelatedSel().All(&ret)

	s.Data["json"] = ret
	s.ServeJSON()

}



type StaffMemberEditController struct {
	beego.Controller
}

// @Title Update
// @Description update the StaffMember
// @Param	Sid		path 	string	true		"The StaffId you want to update"
// @Param	body		body 	models.StaffMember	true		"body for staff content"
// @Success 200 {StaffId} models.StaffMember
// @Failure 403 :StaffId is not int
// @router /:StaffId [put]
func (this *StaffMemberEditController) Put(){
	o := orm.NewOrm()
	StaffId,_ := strconv.Atoi(this.Ctx.Input.Param(":StaffId"))
	var sDetails models.StaffMember

	json.Unmarshal(this.Ctx.Input.RequestBody, &sDetails)

	staffmem := models.StaffMember{StaffId: StaffId}
		if o.Read(&staffmem) == nil {
			
			staffmem.DateOfBirth = sDetails.DateOfBirth
			staffmem.ContactNumber = sDetails.ContactNumber
			staffmem.DateOfJoining = sDetails.DateOfJoining
			staffmem.Gender = sDetails.Gender
			staffmem.BloodGroup = sDetails.BloodGroup
			staffmem.Address = sDetails.Address
			staffmem.Salutation = sDetails.Salutation
			staffmem.ContractType = sDetails.ContractType
			staffmem.LastName = sDetails.LastName
			staffmem.Department = sDetails.Department
			staffmem.FirstName = sDetails.FirstName
			staffmem.Designation =sDetails.Designation
			staffmem.Salary = sDetails.Salary
			staffmem.MiddleName = sDetails.MiddleName
			staffmem.Email = sDetails.Email
			staffmem.EmergencyContactName = sDetails.EmergencyContactName
			staffmem.EmergencyContactNumber = sDetails.EmergencyContactNumber
			staffmem.EmergencyContactRelationship = sDetails.EmergencyContactRelationship
			staffmem.EducationQualification = sDetails.EducationQualification

   			 if num, err := o.Update(&staffmem); err == nil {
					fmt.Println(num)
        			this.Data["json"] = "Updated"
    			}
		}
	this.ServeJSON()	
}


type StaffMemberDeleteController struct {
	beego.Controller
}

// @Title Delete
// @Description delete the StaffMember
// @Param	Sid		path 	string	true		"The StaffId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 StaffId is empty
// @router /:StaffId [delete]
func (this *StaffMemberDeleteController) Delete(){
	o := orm.NewOrm()
	StaffId, _ := strconv.Atoi(this.Ctx.Input.Param(":StaffId"))
			staffD := &models.StaffMember{StaffId:StaffId}
			
			
			if num, err := o.Delete(staffD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type StaffMemberSearchController struct {
	beego.Controller
}

// @Title StaffMemberSearch
// @Description Searching StaffMember
// @Param	body		body 	models.StaffMember	true		"body for staff content"
// @Success 200 {int} models.StaffMember.StaffId
// @Failure 403 body is empty
// @router / [post]
func (r *StaffMemberSearchController) Post(){
	
	var Key models.Keyword


	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &Key); err == nil {
		if ret, success, err := models.SearchStaff(&Key); err == nil {
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

type StaffMemberSController struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get StaffMember
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.StaffMember
// @Failure 403
// @router / [get]
func (c *StaffMemberSController) GetAll() {
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

	l, err := models.GetAllStaffMember(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}