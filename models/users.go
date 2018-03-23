package models

import (
	"errors"
	"fmt"
	"github.com/gogather/com"
	"github.com/nickvellios/gojwt"
	"github.com/astaxie/beego/orm"
)



type Users struct {
	Id                int    `orm:"column(pkid);auto"`
	FkidUserStaffid   uint64 `orm:"column(fkid_user_staffid);null" description:"Staff fkid from staff_member table"`
	FkidUserStudentid uint64 `orm:"column(fkid_user_studentid);null" description:"primark key from student_personal_details"`
	RId               string `orm:"column(rId);null"`
	Username          string `orm:"column(username);null"`
	Designation       string `orm:"column(designation);null"`
	Roles             string `orm:"column(roles);null"`
	Password          string `orm:"column(password);null"`
	UserType          int8   `orm:"column(user_type);null" description:"1- Student2- Staff3- Superadmin; 4-Notice Board"`
	Active            int    `orm:"column(active);null"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

/* 
	check
	{
  		"Username":"akshay",
  		"Password":"hello"
	}
*/

func AuthenticateUser(u *Users) (success string, err error, token string, username string){
	user, err := FindUser(u.Username)
	if(err!=nil) {
		success := "no record found"
		err := errors.New("Wrong username Credentials")
		token := ""
		username := ""
		return success, err, token, username

	} else {
		entered_pass := u.Password
		act_pass := user.Password
		err	:=	VerifyPassword(entered_pass, act_pass, user.RId)
		if err != nil {
			success := "Wrong Password"
			err = errors.New("Wrong Password entered")
			token := ""
			username := "";
			return success, err, token, username

		} else {
			success := "Success!"
			err = nil 
			token:= GenerateToken(user.RId)
			payload, err:=jwt.Decode(token);
			if(err != nil){
				fmt.Print(err)
			} else {
				fmt.Println(payload["Rid"])
			}
			username := user.Username;
			return success, err, token, username

		}
		

	}
}

func FindUser(username string) (Users ,error){
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username:username}
	err := o.Read(&user,"username")
	return user, err
}

func VerifyPassword(entered_pass string, actual_pass string, rId string) (err error){
	Password := com.Md5(entered_pass+rId)
	if(Password != actual_pass){
		return errors.New("wrong password entered")
	}
	return nil
}

func GenerateToken(r_id string) (string){
	rid := r_id
	claims := map[string]string{
		"Rid": rid,
	  }
	expiration := 86400
	gen_token := jwt.Generate(claims,expiration)
	 return gen_token
}

func AddNewUser(u *Users) (string, error){
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.RId = com.RandString(2000)
	user.Username = u.Username 
	user.Designation = u.Designation
	user.Password = com.Md5(u.Password + user.RId)
	o.Insert(user)
	success := "Successfully Inserted"
	return success, nil
}