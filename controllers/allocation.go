package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
)

type AllocationController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create AttandanceAllocation
// @Param	body		body 	models.Allocation	true		"body for Allocation content"
// @Success 201 {int} models.Allocation
// @Failure 403 body is empty
// @router / [post]
func (r *AllocationController) Post(){

	var cd models.Allocation
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddAllocation(&cd); err == nil {
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

type AllocationAllController struct{
	beego.Controller
}

// @Title GetAll
// @Description get all Allocation
// @Success 200 {Allocation} models.Allocation
// @router / [get]
func(dd *AllocationAllController) Get(){
	a := &models.Allocation{}
	d := a.GetAllAllocation()
	dd.Data["json"] = d
	dd.ServeJSON()
}