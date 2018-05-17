package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
	"strconv"
	"AppUpdate/tools/apktools"
	"AppUpdate/models"
	"AppUpdate/tools/fileTools"
	"AppUpdate/tools/tokenTools"
)

type AppController struct {
	beego.Controller
}

func (ctx *AppController) Post(){
	result := make(map[string]interface{})

	token := ctx.Ctx.Input.Header("token")
	beego.Debug("token:  "+token)

	if token == ""{
		result["status"] = "ERROR"
		result["describe"] = "token is nil"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	userEmail, err := tokenTools.CheckToken(token)
	beego.Debug("userEmail:  "+userEmail)
	if err != nil {
		result["status"] = "ERROR"
		result["describe"] = "token is error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	//todo   a

	f, _, err_file := ctx.GetFile("apk")
	if err_file != nil{
		result["status"] = "ERROR"
		result["describe"] = "file is error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}
	fileName := strconv.FormatInt(time.Now().UTC().Unix(),10) + ".apk"

	f.Close()

	todayTime := time.Now().Format("2006-01-02")
	uploadDirPath := "static/uploadfile/"+ todayTime
	isExist := fileTools.PathExist(uploadDirPath)
	if !isExist{
		fileTools.MakeDir(uploadDirPath)
	}

	ctx.SaveToFile("apk", path.Join(uploadDirPath,fileName))

	appName,versionCode,versionName := apktools.GetApkInfo(path.Join(uploadDirPath,fileName))

	intVersionCode, err := strconv.Atoi(versionCode)

	if err != nil{
		intVersionCode = 1
	}

	localAppInfo, apperr := models.GetAppInfo(1,appName,intVersionCode)

	beego.Debug(localAppInfo)
	beego.Debug(apperr)
	if apperr == nil{
		result["status"] = "ERROR"
		result["describe"] = "app already exist"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}
	appInfo := &models.AppModel{}
	appInfo.CompanyId = 1
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
