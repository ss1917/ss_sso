package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:SsoController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:SsoController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:SsoController"] = append(beego.GlobalControllerRouter["github.com/ss1917/ss_sso/controllers:SsoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
