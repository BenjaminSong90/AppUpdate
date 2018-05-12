package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           int
	UserName     string
	Email        string          `orm:"unique"`
	Password     string
	Active       bool			 `orm:"default(true)"`
	Describe	 string
	CreateTime   time.Time       `orm:"auto_now_add;type(datetime)"`
	ModifyTime   time.Time       `orm:"auto_now;type(datetime)"`
}

func (a *User) TableName() string {
	return TableName("user")
}

func UserAdd(role *User) (int64, error) {
	id, err := orm.NewOrm().Insert(role)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func UserDelete(role *User) (num int64, err error) {
	num, err = orm.NewOrm().Delete(role);
	return
}

func GetUserByEmail(email string) ( *User, error) {
	user := User{}
	query := orm.NewOrm().QueryTable(TableName("user"))
	err := query.Filter("email", email).One(&user)
	return &user, err
}
