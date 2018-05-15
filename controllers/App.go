package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"AppUpdate/tools/tokenTools"
	"time"
	"strconv"
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

	f, _, err_file := ctx.GetFile("apk")
	if err_file != nil{
		result["status"] = "ERROR"
		result["describe"] = "file is error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}
	fileName := userEmail+"_"+strconv.FormatInt(time.Now().UTC().Unix(),10)


	f.Close()
	ctx.SaveToFile(fileName, path.Join("static/uploadfile",userEmail))

}
