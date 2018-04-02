package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["CollegeMis/controllers:AllocationAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:AllocationAllController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:AllocationController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:AllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Bid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookEditController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Bid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookSController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookSectionController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookSectionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:BookSectionSController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:BookSectionSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ContractTypeAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ContractTypeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ContractTypeController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ContractTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:CourseAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:CourseAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:CourseController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DeleteAllocationController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DeleteAllocationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Aaid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DepartmentAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DepartmentAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DepartmentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DesignationAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DesignationAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DesignationController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DesignationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DivisionAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DivisionAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DivisionController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:DivisionViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:DivisionViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Cid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeEditController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeSController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeSearchController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeEditController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeTypeViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:NoticeViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:NoticeViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:PractAllocationController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:PractAllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:PracticalAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:PracticalAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:PracticalDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:PracticalDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Prid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:PracticalSAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:PracticalSAllController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:PracticalViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:PracticalViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Did`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ProjAllocationController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ProjAllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ProjectAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ProjectAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ProjectDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ProjectDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Paid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:ProjectSAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:ProjectSAllController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:RegisterController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberEditController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberSController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberSearchController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StaffMemberViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentDeleteController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentEditController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentSController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentSearchController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:StudentViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:StudentViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:SubjectChoiceController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:SubjectChoiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsCAllController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsCAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsViewController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:SubjectsViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Subid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["CollegeMis/controllers:UsersController"] = append(beego.GlobalControllerRouter["CollegeMis/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
