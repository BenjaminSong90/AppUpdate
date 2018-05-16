package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"strconv"
	"AppUpdate/tools/apktools"
	"AppUpdate/models"
)

type AppController struct {
	beego.Controller
}

func (ctx *AppController) Post(){
	result := make(map[string]interface{})
	//
	//token := ctx.Ctx.Input.Header("token")
	//beego.Debug("token:  "+token)
	//
	//if token == ""{
	//	result["status"] = "ERROR"
	//	result["describe"] = "token is nil"
	//	ctx.Data["json"] = result
	//	ctx.ServeJSON()
	//	return
	//}
	//
	//userEmail, err := tokenTools.CheckToken(token)
	//beego.Debug("userEmail:  "+userEmail)
	//if err != nil {
	//	result["status"] = "ERROR"
	//	result["describe"] = "token is error"
	//	ctx.Data["json"] = result
	//	ctx.ServeJSON()
	//	return
	//}
	//
	f, _, err_file := ctx.GetFile("apk")
	if err_file != nil{
		result["status"] = "ERROR"
		result["describe"] = "file is error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}
	fileName := strconv.FormatInt(time.Now().UTC().Unix(),10)


	f.Close()
	todayTime := time.Now().Format("2006-01-02")
	ctx.SaveToFile("apk", path.Join("static/uploadfile",fileName))

	appName,versionCode,versionName := apktools.GetApkInfo(path.Join("static/uploadfile",fileName))

	intVersionCode, err := strconv.Atoi(versionCode)

	if err != nil{
		intVersionCode = 1
	}

	_, apperr := models.GetAppInfo(1,appName,intVersionCode)

	beego.Debug(apperr)
	if apperr == nil{
		result["status"] = "ERROR"
		result["describe"] = "app already exist"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}
	appInfo := &models.AppModel{}
	appInfo.AppName = appName
	appInfo.AppVersionName = versionName
	appInfo.AppVersionNum = intVersionCode
	appInfo.AppSavePath = todayTime+"/"+fileName

	models.AppInfoAdd(appInfo)


	result["status"] = "SUCCESS"
	result["describe"] = "app save success "
	ctx.Data["json"] = result
	ctx.ServeJSON()
	return


}
