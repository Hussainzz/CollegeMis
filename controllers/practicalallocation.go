package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"CollegeMis/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"fmt"
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