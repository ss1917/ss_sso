package models

import (
	"time"
	//"github.com/astaxie/beego/orm"
)

type Role struct {
	RoleId   int       `orm:"column(role_id);auto" json:"role_id"`
	RegionId int       `orm:"column(region_id);null"`
	RoleName     string    `orm:"column(name);size(50);null" json:"rolename"`
	Status   int16     `orm:"column(status);null"`
	Uptime   time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime   time.Time `orm:"column(crtime);type(datetime);null"`
}

func (t *Role) TableName() string {
	return "roles"
}
