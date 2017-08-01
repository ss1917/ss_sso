package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type RoleFunction struct {
	Id         int       `orm:"column(role_function_id);auto"`
	RoleId     int       `orm:"column(role_id);null"`
	FunctionId int       `orm:"column(function_id);null"`
	Status     int16     `orm:"column(status);null"`
	Uptime     time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime     time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *RoleFunction) TableName() string {
	return "role_functions"
}

func GetFunctionIdByRoleId(rid int64) (funclist orm.ParamsList) {
	u := new(RoleFunction)
	o := orm.NewOrm()
	o.Using("read")
	o.QueryTable(u).Filter("Status", 0).Filter("RoleId", rid).ValuesFlat(&funclist, "FunctionId")

	//fmt.Println(funclist)
	return
}