package models

import "time"

type Company struct {
	Id 				int
	CompanyName 	string
	Email 			string 		`orm:"unique"`
	Active			bool		`orm:"default(true)"`
	CreateTime		time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyTime		time.Time 	`orm:"auto_now;type(datetime)"`
}

func (a *Company) TableName() string {
	return TableName("company")
}
