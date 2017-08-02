package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/ss1917/ss_sso/libs/utils"
	"github.com/ss1917/ss_sso/models"
	"strconv"
	"time"
)

// Operations about login
type LoginController struct {
	beego.Controller
}

func (this *LoginController) Prepare() {
}

// @Title get
// @Description get login
// @Success 200 "请求方式有误"
// @Failure 403 body is empty
// @router / [get]
func (this *LoginController) Get() {
	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "请求方式有误",
	}
	this.ServeJSON()
}

// 登录
// swagger:operation POST /v1/accounts/login
// @Title user login
// @Description login
// @Param username        path    string  true        "the login"
// @Success 200 {string}
// @Failure 403 body is empty
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
	fmt.Println(newPass)
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
		fmt.Println(token)
		if autologin {
			this.Ctx.SetCookie("auth_key", token, 604800, "/", Domain)
		} else {
			this.Ctx.SetCookie("auth_key", token, 86400, "/", Domain)
		}

		req := httplib.Put(beego.AppConfig.String("verify_url"))
		req.SetTimeout(1*time.Second, 1*time.Second)
		req.Param("uid", strconv.Itoa(user.UserId))
		req.String()

		this.Data["json"] = map[string]interface{}{
			"status": 0,
			"msg":    "登录成功",
		}
		this.ServeJSON()
	}
}
