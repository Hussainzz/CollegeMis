package models

import (
	"github.com/astaxie/beego/orm"
)

type Designation struct {
	Dsid int `orm:"pk;auto" json:"Dsid"`
	DesignationName string `json:"DesignationName"`
}

type ContractType struct {
	Cid int `orm:"pk; auto" json:"Cid"`
	ContractName string `json:"ContractName"`
}

func init(){
	orm.RegisterModel(new(Designation), new(ContractType))
}

func AddingDesignation(dDetails *Designation) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	designation := &Designation{}
	
	designation.DesignationName = dDetails.DesignationName
	
	o.Insert(designation)
	success = "New Designation Added Successfully"
	return success, nil
}

func AddingContractType(cDetails *ContractType) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	contract := &ContractType{}
	
	contract.ContractName = cDetails.ContractName
	
	o.Insert(contract)
	success = "New ContractType Added Successfully"
	return success, nil
}


func (designation Designation) GetAllDesignation() (ret []Designation) {
	o := orm.NewOrm()
	o.QueryTable("designation").All(&ret)
	return
}

func (contract ContractType) GetAllContract() (ret []ContractType) {
	o := orm.NewOrm()
	o.QueryTable("contract_type").All(&ret)
	return
}