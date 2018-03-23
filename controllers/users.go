package controllers

import (
	"CollegeMis/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type Response struct {
	Message 	string 		`json:"message"`
	ErrorType 	string 		`json:"errorType"`
	Token 		string		`json:"token"`
	Username	string 		`json:"username"`
} 

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description Login
// @Param	body		body 	models.Users	true		"body for user content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *UsersController) Post() {
	var v models.Users
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if success, err, token, username := models.AuthenticateUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			errorType := "null"
			ResponseJson := Response {
				Message 	: success,
				ErrorType   : errorType,
				Token		: token,
				Username 	: username,
			}
			c.Data["json"] = ResponseJson
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

type RegisterController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description Registering Users
// @Param	body		body 	models.Users	true		"body for users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (r *RegisterController) Post() {
	var v models.Users
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &v); err == nil{
		if success, err := models.AddNewUser(&v); err == nil {
			r.Ctx.Output.SetStatus(201)
			errorType := "null"
			ResponseJson := Response {
				Message		: success,
				ErrorType 	: errorType, 
			}
			r.Data["json"] = ResponseJson
		} else {
			r.Data["json"] = err.Error()
		}
	} else {
		r.Data["json"] = err.Error()
	}
	r.ServeJSON()
}