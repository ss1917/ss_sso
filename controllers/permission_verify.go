package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"github.com/ss1917/ss_sso/models"
	"strconv"
	"strings"
)

type PermissionVerifyController struct {
	beego.Controller
}

func (this *PermissionVerifyController) Prepare() {
	//fmt.Println(this.Ctx.Request.Method)
	//fmt.Println(this.Ctx.Request.URL)

}

// @router / [get]
func (this *PermissionVerifyController) Get() {
	this.Data["json"] = map[string]interface{}{
		"status": 0,
		"msg":    "鉴权",
	}
	this.ServeJSON()
}

//从redis读取数据对用户访问路径进行鉴权
// @router / [post]
func (this *PermissionVerifyController) Post() {
	uid := this.GetString("uid")
	uri := this.GetString("uri")
	meth := this.GetString("meth")
	//判断是否是超级管理员
	if id, err := strconv.Atoi(uid); err == nil {
		if models.CheckIsSuperuser(id) {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "is superuser",
			}
			this.ServeJSON()
		}
	}

	conn, err := redis.Dial("tcp", beego.AppConfig.String("redis_default::host"))
	if err != nil {
		logs.Error("连接redis错误", err)
	}

	defer conn.Close()
	conn.Do("auth", beego.AppConfig.String("redis_default::password"))
	conn.Do("SELECT", beego.AppConfig.String("redis_default::db"))

	b := bytes.Buffer{}
	b.WriteString(uid)
	b.WriteString(meth)
	v, _ := conn.Do("SMEMBERS", b.String())
	for _, ve := range v.([]interface{}) {
		if strings.HasPrefix(uri, string(ve.([]byte))) {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "success",
			}
			this.ServeJSON()
		}
	}
	//ALL
	c := bytes.Buffer{}
	c.WriteString(uid)
	c.WriteString("ALL")
	v1, _ := conn.Do("SMEMBERS", c.String())
	for _, ve := range v1.([]interface{}) {
		if strings.HasPrefix(uri, string(ve.([]byte))) {
			this.Data["json"] = map[string]interface{}{
				"status": 0,
				"msg":    "success",
			}
			this.ServeJSON()
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": -1,
		"msg":    "没有权限",
	}
	this.ServeJSON()
}

//登录调用初始化用户权限到redis
// @router / [put]
func (this *PermissionVerifyController) Put() {
	uid := this.GetString("uid")
	var methList = []string{"GET", "POST", "PATCH", "DELETE", "PUT", "ALL"}

	conn, err := redis.Dial("tcp", beego.AppConfig.String("redis_default::host"))
	if err != nil {
		logs.Error("连接redis错误", err)
	}

	defer conn.Close()
	conn.Do("auth", beego.AppConfig.String("redis_default::password"))
	conn.Do("SELECT", beego.AppConfig.String("redis_default::db"))

	//清除用户权限
	for _, meth := range methList {
		b := bytes.Buffer{}
		b.WriteString(uid)
		b.WriteString(meth)
		conn.Do("del", b.String())
	}

	if id, err := strconv.Atoi(uid); err == nil {
		for _, rid := range models.GetRoleIdByUserId(id) {
			for _, fid := range models.GetFunctionIdByRoleId(rid.(int64)) {

				b := bytes.Buffer{}
				b.WriteString(uid)
				b.WriteString(models.GetFunctionIdByFuncId(fid.(int64))[0][1].(string))
				//把用户权限写入redis
				_, err = conn.Do("SADD", b.String(), models.GetFunctionIdByFuncId(fid.(int64))[0][0].(string))

				if err != nil {
					logs.Error("写入权限出错", err)
				}
				logs.Info("写入权限", uid, models.GetFunctionIdByFuncId(fid.(int64))[0])
			}
		}
	}

	this.Data["json"] = map[string]interface{}{
		"status": 0,
		"msg":    "success",
	}
	this.ServeJSON()
}
