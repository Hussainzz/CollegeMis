package controllers

import (
	"CollegeMis/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type NoticeTypeController struct{
	beego.Controller
}

// Post ...
// @Title Post
// @Description create NoticeType
// @Param	body		body 	models.NoticeType	true		"body for NoticeType content"
// @Success 201 {int} models.NoticeType
// @Failure 403 body is empty
// @router / [post]
func (r *NoticeTypeController) Post(){

	var ntd models.NoticeType
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &ntd); err == nil {
		if success, err := models.AddingNoticeType(&ntd); err == nil {
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


type NoticeTypeAllController struct{
	beego.Controller
}

// @Title GetAll
// @Description get all NoticeType
// @Success 200 {object} models.NoticeType
// @router / [get]
func(n *NoticeTypeAllController) Get(){
	noticeT := &models.NoticeType{}
	nTypes := noticeT.GetAllNoticType()
	n.Data["json"] = nTypes
	n.ServeJSON()
}

type NoticeTypeViewController struct{
	beego.Controller
}
// @Title Get
// @Description find NoticeType by Ntid
// @Param	Ntid		path 	string	true		"the Ntid you want to get"
// @Success 200 {object} models.NoticeType
// @Failure 403 :Ntid is empty
// @router /:Ntid [get]
func (n *NoticeTypeViewController) Get(){
	Ntid,_ := strconv.Atoi(n.Ctx.Input.Param(":Ntid"))
	noticetype := &models.NoticeType{Ntid:Ntid}
	noticetype.ReadDB()

	n.Data["json"] = noticetype
	n.ServeJSON()

}

type NoticeTypeEditController struct{
	beego.Controller
}
	// @Title Update
	// @Description update the NoticeType
	// @Param	Ntid		path 	string	true		"The Ntid you want to update"
	// @Param	body		body 	models.NoticeType	true		"body for notice content"
	// @Success 200 {Ntid} models.NoticeType
	// @Failure 403 :Ntid is not int
	// @router /:Ntid [put]
func (this *NoticeTypeEditController) Put(){
	o := orm.NewOrm()
	Ntid,_ := strconv.Atoi(this.Ctx.Input.Param(":Ntid"))
	var ntd models.NoticeType

	json.Unmarshal(this.Ctx.Input.RequestBody, &ntd)

	nn := models.NoticeType{Ntid: Ntid}
		if o.Read(&nn) == nil {
	
			nn.NoticeTypeName = ntd.NoticeTypeName
		
   			 if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
        			this.Data["json"] = "Updated"
    			}
		}
	this.ServeJSON()	
}




type NoticeTypeDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the NoticeType
// @Param	Ntid		path 	string	true		"The Ntid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Ntid is empty
// @router /:Ntid [delete]
func (this *NoticeTypeDeleteController) Delete(){
	o := orm.NewOrm()
	Ntid, _ := strconv.Atoi(this.Ctx.Input.Param(":Ntid"))
			noticeTypeD := &models.NoticeType{Ntid:Ntid}
			
			
			if num, err := o.Delete(noticeTypeD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}