package models

import "time"

type AppTestGroup struct {
	Id 			int
	AppName		string
	CompanyId	int
	UserId		int
	Active		bool		`orm:"default(true)"`
	CreateTime	time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyTime	time.Time 	`orm:"auto_now;type(datetime)"`
}
func (a *AppTestGroup) TableName() string {
	return TableName("test_group")
}