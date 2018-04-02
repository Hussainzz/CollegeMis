package controllers

import (
	"CollegeMis/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/orm"

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

type SubjectChoiceAController struct {
	beego.Controller
}

type SubjectsCAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all SubjectsChoice
// @Success 200 {object} models.StudentSubjectChoices
// @router / [get]
func(nn *SubjectsCAllController) Get(){
	subjectsC := &models.StudentSubjectChoices{}
	n := subjectsC.GetAllSubjectChoices()
	nn.Data["json"] = n
	nn.ServeJSON()
}

