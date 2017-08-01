package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Prepare() {

}

// @router / [get]
func (this *LogoutController) Get() {
	this.Ctx.SetCookie("auth_key", "")
	this.Data["json"] = map[string]interface{}{
		"status": 0,
		"msg":    "登出成功",
	}
	this.ServeJSON()
}
