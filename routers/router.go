// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"CollegeMis/controllers"
	
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Login",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/Register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/studentdetails",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
		//This is that multiSearch
		beego.NSNamespace("/studentGetAll",
			beego.NSInclude(
				&controllers.StudentSController{},
			),
		),
		beego.NSNamespace("/getstudentdetails",
			beego.NSInclude(
				&controllers.StudentAllController{},
			),
		),
		beego.NSNamespace("/student",
			beego.NSInclude(
				&controllers.StudentViewController{},
			),
		),
		beego.NSNamespace("/studentEdit",
			beego.NSInclude(
				&controllers.StudentEditController{},
			),
		),
		beego.NSNamespace("/studentDelete",
			beego.NSInclude(
				&controllers.StudentDeleteController{},
			),
		),
		beego.NSNamespace("/studentSearch",
			beego.NSInclude(
				&controllers.StudentSearchController{},
			),
		),
		beego.NSNamespace("/addcourse",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),
		beego.NSNamespace("/getcourse",
			beego.NSInclude(
				&controllers.CourseAllController{},
			),
		),
		beego.NSNamespace("/adddivision",
			beego.NSInclude(
				&controllers.DivisionController{},
			),
		),
		beego.NSNamespace("/getdivision",
			beego.NSInclude(
				&controllers.DivisionAllController{},
			),
		),

		beego.NSNamespace("/finddivision",
			beego.NSInclude(
				&controllers.DivisionViewController{},
			),
		),

		beego.NSNamespace("/addsubject",
			beego.NSInclude(
				&controllers.SubjectsController{},
			),
		),

		beego.NSNamespace("/addsubjectchoice",
			beego.NSInclude(
				&controllers.SubjectChoiceController{},
			),
		),


		beego.NSNamespace("/staffdetails",
			beego.NSInclude(
				&controllers.StaffMemberController{},
			),
		),
		beego.NSNamespace("/staff",
			beego.NSInclude(
				&controllers.StaffMemberViewController{},
			),
		),
		beego.NSNamespace("/getstaffmember",
			beego.NSInclude(
				&controllers.StaffMemberAllController{},
			),
		),
		beego.NSNamespace("/staffmemberGetAll",
			beego.NSInclude(
				&controllers.StaffMemberSController{},
			),
		),
		beego.NSNamespace("/staffmemberedit",
			beego.NSInclude(
				&controllers.StaffMemberEditController{},
			),
		),
		beego.NSNamespace("/staffmemberdelete",
			beego.NSInclude(
				&controllers.StaffMemberDeleteController{},
			),
		),
		beego.NSNamespace("/addDepartment",
			beego.NSInclude(
				&controllers.DepartmentController{},
			),
		),
		beego.NSNamespace("/addDesignation",
			beego.NSInclude(
				&controllers.DesignationController{},
			),
		),
		beego.NSNamespace("/addContractType",
			beego.NSInclude(
				&controllers.ContractTypeController{},
			),
		),
		beego.NSNamespace("/getContractType",
			beego.NSInclude(
				&controllers.ContractTypeAllController{},
			),
		),
		beego.NSNamespace("/getDesignation",
			beego.NSInclude(
				&controllers.DesignationAllController{},
			),
		),
		beego.NSNamespace("/getDepartment",
			beego.NSInclude(
				&controllers.DepartmentAllController{},
			),
		),

		beego.NSNamespace("/newnotice",
			beego.NSInclude(
				&controllers.NoticeController{},
			),
		),
		beego.NSNamespace("/getnotice",
			beego.NSInclude(
				&controllers.NoticeAllController{},
			),
		),
		beego.NSNamespace("/notice",
			beego.NSInclude(
				&controllers.NoticeViewController{},
			),
		),
		beego.NSNamespace("/noticeedit",
			beego.NSInclude(
				&controllers.NoticeEditController{},
			),
		),
		beego.NSNamespace("/noticedelete",
			beego.NSInclude(
				&controllers.NoticeDeleteController{},
			),
		),
		beego.NSNamespace("/noticeSearch",
			beego.NSInclude(
				&controllers.NoticeSearchController{},
			),
		),
		beego.NSNamespace("/noticeS",
			beego.NSInclude(
				&controllers.NoticeSController{},
			),
		),

		beego.NSNamespace("/newnoticeType",
			beego.NSInclude(
				&controllers.NoticeTypeController{},
			),
		),
		beego.NSNamespace("/getnoticeType",
			beego.NSInclude(
				&controllers.NoticeTypeAllController{},
			),
		),
		beego.NSNamespace("/noticeType",
			beego.NSInclude(
				&controllers.NoticeTypeViewController{},
			),
		),
		beego.NSNamespace("/noticeTypeEdit",
			beego.NSInclude(
				&controllers.NoticeTypeEditController{},
			),
		),
		beego.NSNamespace("/noticeTypeDelete",
			beego.NSInclude(
				&controllers.NoticeTypeDeleteController{},
			),
		),

	
		beego.NSNamespace("/addallocation",
			beego.NSInclude(
				&controllers.AllocationController{},
			),
		),
		beego.NSNamespace("/getallocation",
			beego.NSInclude(
				&controllers.AllocationAllController{},
			),
		),

		beego.NSNamespace("/addprojectallocation",
			beego.NSInclude(
				&controllers.ProjAllocationController{},
			),
		),
		beego.NSNamespace("/getprojectallocation",
			beego.NSInclude(
				&controllers.ProjectAllController{},
			),
		),

		beego.NSNamespace("/addpracticalallocation",
			beego.NSInclude(
				&controllers.PractAllocationController{},
			),
		),
		beego.NSNamespace("/getpracticalallocation",
			beego.NSInclude(
				&controllers.PracticalAllController{},
			),
		),
		beego.NSNamespace("/practical",
			beego.NSInclude(
				&controllers.PracticalViewController{},
			),
		),

		//--------------Library Module----------------
		beego.NSNamespace("/addbooksection",
			beego.NSInclude(
				&controllers.BookSectionController{},
			),
		),
		beego.NSNamespace("/addbook",
			beego.NSInclude(
				&controllers.BookController{},
			),
		),
		beego.NSNamespace("/getallbooks",
			beego.NSInclude(
				&controllers.BookSController{},
			),
		),
		beego.NSNamespace("/editbook",
			beego.NSInclude(
				&controllers.BookEditController{},
			),
		),
		beego.NSNamespace("/deletebook",
			beego.NSInclude(
				&controllers.BookDeleteController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
