package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:AllocationAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:AllocationAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:AllocationController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:AllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookDeleteController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Bid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookEditController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Bid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookSController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookSectionController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:BookSectionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ContractTypeAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ContractTypeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ContractTypeController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ContractTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:CourseAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:CourseAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:CourseController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DepartmentAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DepartmentAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DepartmentController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DepartmentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DesignationAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DesignationAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DesignationController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DesignationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:DivisionViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Cid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeDeleteController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeEditController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeSController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeSearchController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeDeleteController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeEditController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeTypeViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Ntid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:NoticeViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Nid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PractAllocationController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PractAllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PracticalAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PracticalAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PracticalViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:PracticalViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Did`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ProjAllocationController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ProjAllocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ProjectAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:ProjectAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:RegisterController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberDeleteController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberEditController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberSController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberSearchController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StaffMemberViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:StaffId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentDeleteController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentDeleteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentEditController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentEditController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentSController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentSController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentSearchController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentSearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:StudentViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Sid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectChoiceController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectChoiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsAllController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsAllController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsViewController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:SubjectsViewController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:Subid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:UsersController"] = append(beego.GlobalControllerRouter["MainApiWorking/CollegeMis/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
