package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"ss_sso/common/utils"
	"ss_sso/models"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Prepare() {
	this.EnableXSRF = false
}

// @router / [get]
func (this *LoginController) Get() {
	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "请求方式有误",
	}
	this.ServeJSON()
}

// 登录
// @router / [post]
func (this *LoginController) Post() {
	//获取参数
	type UserInfo struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Autologin string `json:"autologin"`
	}
	var user_info *UserInfo = new(UserInfo)
	json.Unmarshal(this.Ctx.Input.RequestBody, user_info)

	autologin := user_info.Autologin == "on"
	//密码加密
	newPass := sjwt.GenMD5(user_info.Password)
	user, err := models.UserGetByName(user_info.Username)
	if err != nil || user.Password != newPass {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "账号密码错误",
		}
		this.ServeJSON()

	} else if user.Status == -10 {
		this.Data["json"] = map[string]interface{}{
			"status": -1,
			"msg":    "该帐号已禁用",
		}
		this.ServeJSON()

	} else {
		user.Last_login = time.Now()
		models.UserUpdate(user)
		token := sjwt.GenToken(user_info.Username, user.UserId, user.RegionId)
		Domain := beego.AppConfig.String("domain")

		if autologin {
			this.Ctx.SetCookie("auth_key", token, 7*86400, "/", Domain)
		} else {
			this.Ctx.SetCookie("auth_key", token, 3600, "/", Domain)
		}

		this.Data["json"] = map[string]interface{}{
			"status": 0,
			"msg":    "登录成功",
		}
		this.ServeJSON()
	}
}
