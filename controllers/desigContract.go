package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
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

type DesignationDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Designation
// @Param	Dsid		path 	string	true		"The Dsid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Dsid is empty
// @router /:Dsid [delete]
func (this *DesignationDeleteController) Delete(){
	o := orm.NewOrm()
	Dsid, _ := strconv.Atoi(this.Ctx.Input.Param(":Dsid"))
			DesignationD := &models.Designation{Dsid:Dsid}
			
			
			if num, err := o.Delete(DesignationD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type DesignationEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the Designation
// @Param	Dsid		path 	string	true		"The Dsid you want to update"
// @Param	body		body 	models.Designation	true		"body for Designation content"
// @Success 200 {Dsid} models.Designation
// @Failure 403 :Dsid is not int
// @router /:Dsid [put]
func (this *DesignationEditController) Put(){
	o := orm.NewOrm()
	Dsid,_ := strconv.Atoi(this.Ctx.Input.Param(":Dsid"))
	var nd models.Designation

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.Designation{Dsid: Dsid}
		if o.Read(&nn) == nil {
	
			nn.DesignationName = nd.DesignationName
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
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

type ContractDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the ContractType
// @Param	Cid		path 	string	true		"The Cid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Cid is empty
// @router /:Cid [delete]
func (this *ContractDeleteController) Delete(){
	o := orm.NewOrm()
	Cid, _ := strconv.Atoi(this.Ctx.Input.Param(":Cid"))
			ContractTypeD := &models.ContractType{Cid:Cid}
			
			
			if num, err := o.Delete(ContractTypeD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type ContractEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the ContractType
// @Param	Cid		path 	string	true		"The Cid you want to update"
// @Param	body		body 	models.ContractType	true		"body for BookSection content"
// @Success 200 {Cid} models.ContractType
// @Failure 403 :Cid is not int
// @router /:Cid [put]
func (this *ContractEditController) Put(){
	o := orm.NewOrm()
	Cid,_ := strconv.Atoi(this.Ctx.Input.Param(":Cid"))
	var nd models.ContractType

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.ContractType{Cid: Cid}
		if o.Read(&nn) == nil {
	
			nn.ContractName = nd.ContractName
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
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

type DepartmentDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Department
// @Param	Depid		path 	string	true		"The Depid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Depid is empty
// @router /:Depid [delete]
func (this *DepartmentDeleteController) Delete(){
	o := orm.NewOrm()
	Depid, _ := strconv.Atoi(this.Ctx.Input.Param(":Depid"))
			DepartmentD := &models.Department{Depid:Depid}
			
			
			if num, err := o.Delete(DepartmentD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type DepartmentEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the 
// @Param	Bsid		path 	string	true		"The Depid you want to update"
// @Param	body		body 	models.Department	true		"body for Department content"
// @Success 200 {Depid} models.Department
// @Failure 403 :Depid is not int
// @router /:Depid [put]
func (this *DepartmentEditController) Put(){
	o := orm.NewOrm()
	Depid,_ := strconv.Atoi(this.Ctx.Input.Param(":Depid"))
	var nd models.Department

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.Department{Depid: Depid}
		if o.Read(&nn) == nil {
	
			nn.DepartmentName = nd.DepartmentName
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
}