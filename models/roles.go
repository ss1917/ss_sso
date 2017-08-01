package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
)

var (
	Roles map[string]*Role
)

//status 20：逻辑删除，10：禁用， 0：正常
type Role struct {
	RoleId   int       `orm:"column(role_id);auto" json:"role_id"`
	RoleName string    `orm:"column(name);size(50);null" json:"rolename"`
	Status   int16     `orm:"column(status);null"`
	Uptime   time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime   time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *Role) TableName() string {
	return "roles"
}

func RoleGetAll() (maps []orm.Params, err error) {
	o := orm.NewOrm()
	o.Using("user")
	u := new(Role)
	if _, err = o.QueryTable(u).Exclude("Status", 20).Values(&maps, "RoleId",
					"RegionId", "RoleName", "Status", "Uptime"); err != nil {
		return
	}
	return
}
