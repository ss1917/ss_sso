package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

var (
	METHOD_TYPE_ALL    = 0  // ALL
	METHOD_TYPE_GET    = 10 // GET
	METHOD_TYPE_POST   = 20 // POST
	METHOD_TYPE_PUT    = 30 // PUT
	METHOD_TYPE_PATCH  = 40 // PATCH
	METHOD_TYPE_DELETE = 50 // DELETE
)
var (
	Functions map[string]*Function
)

type Function struct {
	Id         int       `orm:"column(function_id);auto"`
	Name       string    `orm:"column(name);size(60);null"`
	Uri        string    `orm:"column(uri);size(300);null"`
	MethodType string     `orm:"column(method_type);null"`
	MenuName   string    `orm:"column(menuname);size(150);null"`
	Icon       string    `orm:"column(icon);size(20);null"`
	Level      string    `orm:"column(level);size(3);null"`
	Sort       string    `orm:"column(sort);size(3);null"`
	Status     int16     `orm:"column(status);null"`
	Uptime     time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime     time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *Function) TableName() string {
	return "functions"
}

func GetFunctionIdByFuncId(fid int64) (funclist []orm.ParamsList) {
	u := new(Function)
	o := orm.NewOrm()
	o.Using("read")
	o.QueryTable(u).Filter("Status", 0).Filter("Id", fid).ValuesList(&funclist, "Uri","MethodType")

	//fmt.Println(funclist)
	return
}