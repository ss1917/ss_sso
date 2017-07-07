package models

import (
	//"github.com/astaxie/beego/orm"
	"time"
)

type UserRole struct {
	UserRoleId int       `orm:"column(user_role_id);auto" json:"user_role_id"`
	UserId     int       `orm:"column(user_id);null" json:"user_id"`
	RoleId     int       `orm:"column(role_id);null" json:"role_id"`
	RegionId   int       `orm:"column(region_id);null" json:"region_id"`
	Status     int16     `orm:"column(status);null" json:"status"`
	Uptime     time.Time `orm:"column(uptime);type(datetime);null" json:"uptime"`
	Crtime     time.Time `orm:"column(crtime);type(datetime);null" json:"crtime"`
}

func (t *UserRole) TableName() string {
	return "user_roles"
}
