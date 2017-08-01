package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/ss1917/ss_sso/libs/utils"
)

type SsoController struct {
	beego.Controller
}

func (this *SsoController) Prepare() {

}

// @router / [get]
func (this *SsoController) Get() {
	this.Ctx.SetCookie("auth_key", "")
	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "请求方式有误",
	}
	this.ServeJSON()
}

// @router / [post]
func (this *SsoController) Post() {
	//fmt.Println(this.Ctx.GetCookie("auth_key"))
	if auth_token := this.GetString("auth_key"); auth_token != "" {
		if username, userid, regionid := sjwt.CheckToken(auth_token); username != "" {
			this.Data["json"] = map[string]interface{}{
				"status":   0,
				"username": username,
				"user_id":  userid,
				"regionid": regionid,
			}
			this.ServeJSON()
		}
	}
	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "验证失败",
	}
	this.ServeJSON()
}
