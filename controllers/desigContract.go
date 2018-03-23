package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
)

type DesignationController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create Designation
// @Param	body		body 	models.Designation	true		"body for Designation content"
// @Success 201 {int} models.Designation
// @Failure 403 body is empty
// @router / [post]
func (r *DesignationController) Post(){

	var cd models.Designation
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddingDesignation(&cd); err == nil {
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

type DesignationAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all Designation
// @Success 200 {object} models.Designation
// @router / [get]
func(cc *DesignationAllController) Get(){
	designation := &models.Designation{}
	c := designation.GetAllDesignation()
	cc.Data["json"] = c
	cc.ServeJSON()
}



type ContractTypeController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create ContractType
// @Param	body		body 	models.ContractType	true		"body for Contract content"
// @Success 201 {int} models.ContractType
// @Failure 403 body is empty
// @router / [post]
func (r *ContractTypeController) Post(){

	var cd models.ContractType
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddingContractType(&cd); err == nil {
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



type ContractTypeAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all ContractType
// @Success 200 {object} models.ContractType
// @router / [get]
func(cc *ContractTypeAllController) Get(){
	contract := &models.ContractType{}
	c := contract.GetAllContract()
	cc.Data["json"] = c
	cc.ServeJSON()
}




type DepartmentController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create Department
// @Param	body		body 	models.Department	true		"body for Department content"
// @Success 201 {int} models.Department
// @Failure 403 body is empty
// @router / [post]
func (r *DepartmentController) Post(){

	var cd models.Department
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddingDepartment(&cd); err == nil {
			r.Ctx.Output.SetStatus(201) 
			//errorType := "null"
		//	ResponseJson := Response {
		//		Message 	: success,
		//		ErrorType   : errorType,
				
			//}
			r.Data["json"] = success
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()
}




type DepartmentAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all Department
// @Success 200 {object} models.Department
// @router / [get]
func(cc *DepartmentAllController) Get(){
	dept := &models.Department{}
	c := dept.GetAllDepartment()
	cc.Data["json"] = c
	cc.ServeJSON()
}