package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AppModel struct {
	Id 				int
	CompanyId		int
	AppName			string
	AppVersionName	string
	AppVersionNum	int
	AppSavePath		string
	Active			bool 		`orm:"default(true)"`
	CreateTime		time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyTime		time.Time 	`orm:"auto_now;type(datetime)"`
}

func (a *AppModel) TableName() string {
	return TableName("app")
}

func GetAppInfo(companyId int, appName string,versionCode int)(*AppModel, error){
	appInfo := &AppModel{}
	appInfo.CompanyId = companyId
	appInfo.AppVersionName = appName
	appInfo.AppVersionNum = versionCode
	query := orm.NewOrm().QueryTable(TableName("app"))
	err := query.Filter("CompanyId", companyId).Filter("AppName", appName).Filter("AppVersionNum",versionCode).One(appInfo)
	return appInfo, err
}


func AppInfoAdd(appInfo *AppModel) (int64, error) {
	id, err := orm.NewOrm().Insert(appInfo)
	if err != nil {
		return -1, err
	}
	return id, nil
}



