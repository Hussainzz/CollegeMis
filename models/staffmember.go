package models

import (
"github.com/astaxie/beego/orm"
	"time"
	"math/rand"
	"reflect"
	"strings"
	"errors"
)

type StaffMember struct {
	StaffId                      int    `orm:"pk;auto" json:"StaffId"`
	Active                       int    `orm:"null" json:"Active"`
	DateOfBirth                  string `orm:"null" json:"DateOfBirth"`
	ContactNumber                int64  `orm:"null" json:"ContactNumber"`
	DateOfJoining                string `orm:"null" json:"DateOfJoining"`
	StaffCode                    string `orm:"null" json:"StaffCode"`
	Gender                       string `orm:"null" json:"Gender"`
	BloodGroup                   string `orm:"null" json:"BloodGroup"`
	Address                      string `orm:"size(700);null" json:"Address"`
	Salutation                   string `orm:"null" json:"Salutation"`
	Userid                       string `orm:"null" json:"Userid"`
	ContractType                 *ContractType `orm:"rel(fk);null" json:"ContractType"`
	LastName                     string `orm:"null" json:"LastName"`
	Department                   *Department `orm:"rel(fk);null" json:"Department"`
	FirstName                    string `orm:"null" json:"FirstName"`
	Designation                  *Designation `orm:"rel(fk);null" json:"Designation"`
	Salary                       string `orm:"null" json:"Salary"`
	MiddleName                   string `orm:"null" json:"MiddleName"`
	Email                        string `orm:"null" json:"Email"`
	EmergencyContactName         string `orm:"null" json:"EmergencyContactName"`
	EmergencyContactNumber       string `orm:"null" json:"EmergencyContactNumber"`
	EmergencyContactRelationship string `orm:"null" json:"EmergencyContactRelationship"`
	Status                       int8   `json:"Status"`
	EducationQualification       string `json:"EducationQualification"`
}


func init() {
	orm.RegisterModel(new(StaffMember))
}

func AddingStaffMember(sDetails *StaffMember) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	var success string
	staffmem := &StaffMember{}
	
	staffmem.Active = 1
	staffmem.DateOfBirth = sDetails.DateOfBirth
	staffmem.ContactNumber = sDetails.ContactNumber
	staffmem.DateOfJoining = sDetails.DateOfJoining
	
	rand.Seed(time.Now().UTC().UnixNano())
	staffmem.StaffCode = "STAFF-"+randomString(4)
	
	staffmem.Gender = sDetails.Gender
	staffmem.BloodGroup = sDetails.BloodGroup
	staffmem.Address = sDetails.Address
	staffmem.Salutation = sDetails.Salutation
	staffmem.ContractType = sDetails.ContractType
	staffmem.LastName = sDetails.LastName
	staffmem.Department = sDetails.Department
	staffmem.FirstName = sDetails.FirstName
	staffmem.Designation =sDetails.Designation
	staffmem.Salary = sDetails.Salary
	staffmem.MiddleName = sDetails.MiddleName
	staffmem.Email = sDetails.Email
	staffmem.EmergencyContactName = sDetails.EmergencyContactName
	staffmem.EmergencyContactNumber = sDetails.EmergencyContactNumber
	staffmem.EmergencyContactRelationship = sDetails.EmergencyContactRelationship
	staffmem.Status = 1
	staffmem.EducationQualification = sDetails.EducationQualification

	o.Insert(staffmem)
	success = "ThankYou For Adding Staff - StaffCode is :- " + staffmem.StaffCode 
	return success, nil
}


func randomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
        bytes[i] = byte(randInt(65, 90))
    }
    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}


func (staffdetails StaffMember) GetAllStaffDetails() (ret []StaffMember) {
	o := orm.NewOrm()
	o.QueryTable("staff_member").RelatedSel().OrderBy("-staff_id").All(&ret)
	return
}

func SearchStaff(key *Keyword) ( []StaffMember, string, error){
	
	o := orm.NewOrm()
	
	query := o.QueryTable("staff_member")

	var ret []StaffMember

	query.Filter("first_name", key.Key).RelatedSel().All(&ret)
	query.Filter("designation_id", key.Key).RelatedSel().All(&ret)
	query.Filter("contract_type_id", key.Key).RelatedSel().All(&ret)
	query.Filter("staff_code",key.Key).RelatedSel().All(&ret)

		success := "Queryed Students"
		return  ret, success, nil
}

func GetAllStaffMember(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(StaffMember))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []StaffMember
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}