package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
	"strings"
	"errors"
)



type PractAllocationController struct{
	beego.Controller
}
// Post ...
// @Title Post
// @Description create PracticalAllocation
// @Param	body		body 	models.PracticalAllocation	true		"body for PracticalAllocation content"
// @Success 201 {int} models.PracticalAllocation
// @Failure 403 body is empty
// @router / [post]
func (r *PractAllocationController) Post(){

	var cd models.PracticalAllocation
	if err := json.Unmarshal(r.Ctx.Input.RequestBody, &cd); err == nil {
		if success, err := models.Practicalallocation(&cd); err == nil {
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

type PracticalViewController struct{
	beego.Controller
	}
// @Title Get
// @Description find RollNumber of Student by Did
// @Param	Did		path 	string	true		"the Did you want to get"
// @Success 200 {object} models.StudentDetails
// @Failure 403 :Did is empty
// @router /:Did [get]
func (s *PracticalViewController) Get(){
	Did,_ := strconv.Atoi(s.Ctx.Input.Param(":Did"))
	var ret []models.StudentDetails
	o := orm.NewOrm()
			
	num, err := o.QueryTable("student_details").Filter("division_id", Did).RelatedSel().All(&ret)
		if err == nil {
    		fmt.Printf("%d posts read\n", num)
    		for _, post := range ret {
       			 fmt.Println(post.RollNo)
   				 }
		}	
		
	s.Data["json"] = ret
	s.ServeJSON()
}



type PracticalAllController struct{
	beego.Controller
}

// @Title GetAll
// @Description get all PracticalAllocation
// @Success 200 {PracticalAllocation} models.PracticalAllocation
// @router / [get]
func(dd *PracticalAllController) Get(){
	a := &models.PracticalAllocation{}
	d := a.GetAllPractAllocation()
	fmt.Println(d)
	dd.Data["json"] = d
	dd.ServeJSON()
}

type PracticalDeleteController struct{
	beego.Controller
}

// @Title Delete
// @Description delete the practicalallocation
// @Param	Prid		path 	string	true		"The Prid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Prid is empty
// @router /:Prid [delete]
func (this *PracticalDeleteController) Delete(){
	o := orm.NewOrm()
	Prid, _ := strconv.Atoi(this.Ctx.Input.Param(":Prid"))
			PracticalD := &models.PracticalAllocation{Prid:Prid}
			
			
			if num, err := o.Delete(PracticalD); err == nil {
				fmt.Println(num)
				fmt.Println(num)
				this.Data["json"] = "Deleted"
			}
	
		this.ServeJSON()

}


type PracticalSAllController struct{
	beego.Controller
}
// GetAll ...
// @Title Get All
// @Description get PracticalAllocation
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.PracticalAllocation
// @Failure 403
// @router / [get]
func (c *PracticalSAllController) GetAll() {
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

	l, err := models.GetAllPracticalAllocation(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}