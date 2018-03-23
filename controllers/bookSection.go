package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
)

type BookSectionController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create BookSection
// @Param	body		body 	models.BookSection	true		"body for BookSection content"
// @Success 201 {int} models.BookSection
// @Failure 403 body is empty
// @router / [post]
func (r *BookSectionController) Post(){

	var cd models.BookSection
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.AddBookSection(&cd); err == nil {
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