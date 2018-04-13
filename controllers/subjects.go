package controllers

import (
	"CollegeMis/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/orm"
	"strings"
	"errors"
	"fmt"

)


type SubjectsController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create Subjects
// @Param	body		body 	models.Subjects	true		"body for Subject content"
// @Success 201 {int} models.Subjects
// @Failure 403 body is empty
// @router / [post]
func (r *SubjectsController) Post(){
	var sd models.Subjects
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &sd); err == nil {
		if success, err := models.AddingSubjects(&sd); err == nil {
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


type SubjectsViewController struct{
	beego.Controller
}
// @Title Get
// @Description find Subjects by Subid
// @Param	Subid		path 	string	true		"the Subid you want to get"
// @Success 200 {object} models.Subjects
// @Failure 403 :Subid is empty
// @router /:Subid [get]
func (n *SubjectsViewController) Get(){
	Subid,_ := strconv.Atoi(n.Ctx.Input.Param(":Subid"))
	
	var ret []models.Subjects
		o := orm.NewOrm()
		
		o.QueryTable("subjects").Filter("Subid__exact", Subid).RelatedSel().All(&ret)
	
		n.Data["json"] = ret
		n.ServeJSON()
}


type SubjectsAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all Subjects
// @Success 200 {object} models.Subjects
// @router / [get]
func(nn *SubjectsAllController) Get(){
	subjects := &models.Subjects{}
	n := subjects.GetAllSubjects()
	nn.Data["json"] = n
	nn.ServeJSON()
}


type SubjectChoiceController struct {
	beego.Controller
}
// Post ...
// @Title Post
// @Description create SubjectChoices of Students
// @Param	body		body 	models.StudentSubjectChoices	true		"body for StudentSubjectChoices content"
// @Success 201 {int} models.StudentSubjectChoices
// @Failure 403 body is empty
// @router / [post]
func (r *SubjectChoiceController) Post(){
	var sd models.StudentSubjectChoices
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &sd); err == nil {
		if success, err := models.AddingStudentSubjectChoice(&sd); err == nil {
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

type SubjectChoiceAllController struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get SubjectChoices
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.StudentSubjectChoices
// @Failure 403
// @router / [get]
func (c *SubjectChoiceAllController) GetAll() {
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

	l, err := models.GetAllSubjectChoices(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

type SubjectDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Subjects
// @Param	Subid		path 	string	true		"The Subid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Subid is empty
// @router /:Subid [delete]
func (this *SubjectDeleteController) Delete(){
	o := orm.NewOrm()
	Subid, _ := strconv.Atoi(this.Ctx.Input.Param(":Subid"))
			SubjectsD := &models.Subjects{Subid:Subid}
			
			
			if num, err := o.Delete(SubjectsD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type SubjectEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the 
// @Param	Subid		path 	string	true		"The Subid you want to update"
// @Param	body		body 	models.Subid	true		"body for Department content"
// @Success 200 {Subid} models.Subid
// @Failure 403 :Subid is not int
// @router /:Subid [put]
func (this *SubjectEditController) Put(){
	o := orm.NewOrm()
	Subid,_ := strconv.Atoi(this.Ctx.Input.Param(":Subid"))
	var nd models.Subjects

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.Subjects{Subid: Subid}
		if o.Read(&nn) == nil {
	
			nn.SubjectCode = nd.SubjectCode
			nn.SubjectName = nd.SubjectName
			nn.Course = nd.Course
			nn.Semester = nd.Semester
			nn.Compulsory = nd.Compulsory
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
}

type SubjectChoiceDeleteController struct{
	beego.Controller
}

// @Title Delete
// @Description delete the SubjectChoice
// @Param	Scid		path 	string	true		"The Scid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Scid is empty
// @router /:Scid [delete]
func (this *SubjectChoiceDeleteController) Delete(){
	o := orm.NewOrm()
	Scid, _ := strconv.Atoi(this.Ctx.Input.Param(":Scid"))
			SubjectChoicesD := &models.StudentSubjectChoices{Scid:Scid}
			
			
			if num, err := o.Delete(SubjectChoicesD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

