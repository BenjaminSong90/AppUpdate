package models

import "time"

type AppModel struct {
	Id 				int
	CompanyId		int
	AppName			string
	AppVersionName	string
	AppVersionNum	int
	Active			bool 		`orm:"default(true)"`
	CreateTime		time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyTime		time.Time 	`orm:"auto_now;type(datetime)"`
}

func (a *AppModel) TableName() string {
	return TableName("app")
}


