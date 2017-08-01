package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type UserRole struct {
	UserRoleId int       `orm:"column(user_role_id);auto" json:"user_role_id"`
	UserId     int       `orm:"column(user_id);null" json:"user_id"`
	RoleId     int       `orm:"column(role_id);null" json:"role_id"`
	Status     int16     `orm:"column(status);null" json:"status"`
	Uptime     time.Time `orm:"column(uptime);type(datetime);null" json:"uptime"`
	Crtime     time.Time `orm:"column(crtime);type(datetime);null" json:"crtime"`
}

func (t *UserRole) TableName() string {
	return "user_roles"
}

func GetRoleIdByUserId(uid int) (rolelist orm.ParamsList) {
	u := new(UserRole)
	o := orm.NewOrm()
	o.Using("read")
	o.QueryTable(u).Filter("Status", 0).Filter("UserId", uid).ValuesFlat(&rolelist, "RoleId")
	return
}
