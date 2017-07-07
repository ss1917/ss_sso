package models

import (
	"time"
)

type RoleFunction struct {
	Id         int       `orm:"column(role_function_id);auto"`
	RoleId     int       `orm:"column(role_id);null"`
	FunctionId int       `orm:"column(function_id);null"`
	RegionId   int       `orm:"column(region_id);null"`
	Status     int16     `orm:"column(status);null"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(datetime);null"`
	CreatedAt  time.Time `orm:"column(created_at);type(datetime);null"`
}

func (t *RoleFunction) TableName() string {
	return "role_function"
}
