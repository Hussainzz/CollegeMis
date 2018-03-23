package controllers

import (
	"CollegeMis/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strings"
	"errors"
)




type NoticeController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create Notice
// @Param	body		body 	models.Notice	true		"body for Notice content"
// @Success 201 {int} models.Notice
// @Failure 403 body is empty
// @router / [post]
func (r *NoticeController) Post(){
	var nd models.Notice
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &nd); err == nil {
		if success, err := models.AddingNotice(&nd); err == nil {
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



type NoticeAllController struct{
	beego.Controller
}
// @Title GetAll
// @Description get all Notice
// @Success 200 {object} models.Notice
// @router / [get]
func(nn *NoticeAllController) Get(){
	notice := &models.Notice{}
	n := notice.GetAllNotice()
	nn.Data["json"] = n
	nn.ServeJSON()
}


type NoticeViewController struct{
	beego.Controller
}
// @Title Get
// @Description find Notice by Nid
// @Param	Nid		path 	string	true		"the Nid you want to get"
// @Success 200 {object} models.Notice
// @Failure 403 :Nid is empty
// @router /:Nid [get]
func (n *NoticeViewController) Get(){
	Nid,_ := strconv.Atoi(n.Ctx.Input.Param(":Nid"))
	
	var ret []models.Notice
		o := orm.NewOrm()
		
		o.QueryTable("notice").Filter("nid__exact", Nid).RelatedSel().All(&ret)
	
		n.Data["json"] = ret
		n.ServeJSON()
}

type NoticeEditController struct{
		beego.Controller
}
	// @Title Update
	// @Description update the Notice
	// @Param	Nid		path 	string	true		"The Nid you want to update"
	// @Param	body		body 	models.Notice	true		"body for notice content"
	// @Success 200 {Nid} models.Notice
	// @Failure 403 :Nid is not int
	// @router /:Nid [put]
	func (this *NoticeEditController) Put(){
		o := orm.NewOrm()
		Nid,_ := strconv.Atoi(this.Ctx.Input.Param(":Nid"))
		var nd models.Notice
	
		json.Unmarshal(this.Ctx.Input.RequestBody, &nd)
	
		nn := models.Notice{Nid: Nid}
			if o.Read(&nn) == nil {
		
				nn.NoticeTitle = nd.NoticeTitle
				nn.NoticeType = nd.NoticeType
				nn.NoticeContent = nd.NoticeContent
	
					if num, err := o.Update(&nn); err == nil {
						fmt.Println(num)
						this.Data["json"] = "Updated"
					}
			}
		this.ServeJSON()	
}


type NoticeDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the Notice
// @Param	Nid		path 	string	true		"The Nid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Nid is empty
// @router /:Nid [delete]
func (this *NoticeDeleteController) Delete(){
	o := orm.NewOrm()
	Nid, _ := strconv.Atoi(this.Ctx.Input.Param(":Nid"))
			noticeD := &models.Notice{Nid:Nid}
			
			
			if num, err := o.Delete(noticeD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}


type NoticeSearchController struct{
	beego.Controller
}

// @Title NoticeSearch
// @Description Searching Notice
// @Param	body		body 	models.Notice	true		"body for notice content"
// @Success 200 {int} models.Notice.Nid
// @Failure 403 body is empty
// @router / [post]
func (r *NoticeSearchController) Post(){

	var Key models.Keyword
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &Key); err == nil {
		if ret, success, err := models.Search(&Key); err == nil {
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




type NoticeSController struct{
	beego.Controller
}
// GetAll ...
// @Title Get All
// @Description get Notice
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Notice
// @Failure 403
// @router / [get]
func (c *NoticeSController) GetAll() {
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

	l, err := models.GetAllNotice(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}