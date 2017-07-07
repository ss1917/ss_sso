package models

import (
	"time"
)

var (
	METHOD_TYPE_GET    = 10 // GET
	METHOD_TYPE_POST   = 20 // POST
	METHOD_TYPE_PUT    = 30 // PUT
	METHOD_TYPE_PATCH  = 40 // PATCH
	METHOD_TYPE_DELETE = 50 // DELETE
)

type Functions struct {
	Id                 int       `orm:"column(function_id);auto"`
	Name               string    `orm:"column(name);size(60);null"`
	RegionId           int       `orm:"column(region_id);null"`
	Uri                string    `orm:"column(uri);size(500);null"`
	MethodType         int16     `orm:"column(method_type);null"`
	MenuName           string     `orm:"column(menuname);size(250);null"`
	Icon              string     `orm:"column(icon);size(20);null"`
	Level              string     `orm:"column(level);size(3);null"`
	Sort              string     `orm:"column(sort);size(3);null"`
	Status             int16     `orm:"column(status);null"`
	Uptime             time.Time `orm:"column(uptime);type(datetime);null"`
	Crtime             time.Time `orm:"column(crtime);type(datetime);null"`
}
