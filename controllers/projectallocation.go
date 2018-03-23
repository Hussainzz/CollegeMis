package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
)




type ProjAllocationController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create ProjectAllocation
// @Param	body		body 	models.ProjectAllocation	true		"body for ProjectAllocation content"
// @Success 201 {int} models.ProjectAllocation
// @Failure 403 body is empty
// @router / [post]
func (r *ProjAllocationController) Post(){

	var cd models.ProjectAllocation
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.Projectallocation(&cd); err == nil {
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

type ProjectAllController struct{
	beego.Controller
}

// @Title GetAll
// @Description get all ProjectAllocation
// @Success 200 {ProjectAllocation} models.ProjectAllocation
// @router / [get]
func(dd *ProjectAllController) Get(){
	a := &models.ProjectAllocation{}
	d := a.GetAllProjAllocation()
	dd.Data["json"] = d
	dd.ServeJSON()
}