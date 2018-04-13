package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
	"strings"
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
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

type BookSectionSController struct{
	beego.Controller
}
// GetAll ...
// @Title Get All
// @Description get Book
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Book
// @Failure 403
// @router / [get]
func (c *BookSectionSController) GetAll() {
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

	l, err := models.GetAllBookSections(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

type BookSectionDeleteController struct{
	beego.Controller
}
// @Title Delete
// @Description delete the BookSection
// @Param	Bsid		path 	string	true		"The Bsid you want to get the BookSection"
// @Success 200 {string} delete success!
// @Failure 403 Bsid is empty
// @router /:Bsid [delete]
func (this *BookSectionDeleteController) Delete(){
	o := orm.NewOrm()
	Bsid, _ := strconv.Atoi(this.Ctx.Input.Param(":Bsid"))
			bookSectionD := &models.BookSection{Bsid:Bsid}
			
			
			if num, err := o.Delete(bookSectionD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}

type BookSectionEditController struct{
	beego.Controller
}
// @Title Update
// @Description update the BookSection
// @Param	Bsid		path 	string	true		"The Bsid you want to update"
// @Param	body		body 	models.BookSection	true		"body for BookSection content"
// @Success 200 {Bsid} models.BookSection
// @Failure 403 :Bsid is not int
// @router /:Bsid [put]
func (this *BookSectionEditController) Put(){
	o := orm.NewOrm()
	Bsid,_ := strconv.Atoi(this.Ctx.Input.Param(":Bsid"))
	var nd models.BookSection

	json.Unmarshal(this.Ctx.Input.RequestBody, &nd)

	nn := models.BookSection{Bsid: Bsid}
		if o.Read(&nn) == nil {
	
			nn.SectionName = nd.SectionName
				if num, err := o.Update(&nn); err == nil {
					fmt.Println(num)
					this.Data["json"] = "Updated"
				}
		}
	this.ServeJSON()	
}