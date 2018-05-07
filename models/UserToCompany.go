package models

import "time"

const (
	Manager = 0
	Normal	= 1
)


type UserToCompany struct {
	Id 			int
	CompanyId	int
	UserId		int
	Active		bool		`orm:"default(true)"`
	Actor		int  		`orm:"default(1)"`
	CreateTime	time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyTime	time.Time 	`orm:"auto_now;type(datetime)"`
}
