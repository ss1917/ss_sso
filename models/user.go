package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	Users map[string]*User
)

type User struct {
	UserId     int       `orm:"column(user_id);auto"`
	RegionId   int       `orm:"column(region_id);null"`
	Username   string    `orm:"column(username);size(120)" json:"name`
	Password   string    `orm:"column(password);size(200)"`
	Email      string    `orm:"column(email);size(200)"`
	Tel        string    `orm:"column(tel);size(11)"`
	Status     int16     `orm:"column(status);null"`
	Superuser  string    `orm:"column(superuser);size(2);"`
	Nickname   string    `orm:"column(nickname);size(64);"`
	Sex        string    `orm:"column(sex);size(10);null"`
	Role       string    `orm:"column(role);size(200);null"`
	LastIp     string    `orm:"column(lastip);size(20);null"`
	Last_login time.Time `orm:"column(last_login);type(datetime);null"`
	Wtime      time.Time `orm:"column(wtime);type(datetime);null"`
}

func (t *User) TableName() string {
	return "user"
}

func UserGetById(id int) (*User, error) {
	u := new(User)

	err := orm.NewOrm().QueryTable(u).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserGetByName(username string) (*User, error) {
	u := new(User)
	o := orm.NewOrm()
	o.Using("read")

	err := o.QueryTable(u).Filter("Username", username).One(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func UserUpdate(user *User, fields ...string) error {
	_, err := orm.NewOrm().Update(user, fields...)
	return err
}
