// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/ss1917/ss_sso/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/accounts",
			beego.NSNamespace("/login",
				beego.NSInclude(
					&controllers.LoginController{},
				),
			),
			beego.NSNamespace("/logout",
				beego.NSInclude(
					&controllers.LogoutController{},
				),
			),
			beego.NSNamespace("/sso",
				beego.NSInclude(
					&controllers.SsoController{},
				),
			),
		),

		beego.NSNamespace("/rbac",
			beego.NSNamespace("/verify",
				beego.NSInclude(
					&controllers.PermissionVerifyController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
