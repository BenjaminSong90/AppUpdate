package controllers

import (
	"github.com/astaxie/beego"
	"AppUpdate/models"
	"AppUpdate/tools/tokenTools"
)

type LoginController struct {
	beego.Controller
}

func (ctx *LoginController) Post() {
	result := make(map[string]interface{})
	email := ctx.Input().Get("email")
	password := ctx.Input().Get("password")

	if email == "" || password ==""{
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "password or email is null"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	user, err := models.GetUserByEmail(email)

	if err != nil {
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "The user is null"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	beego.Debug("<>"+user.Password+"<>"+password)

	if password != user.Password {
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "The password id invalid"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}



	token := tokenTools.CreateToken(user.Email)

	result["status"] = "SUCCESS"
	result["token"] = token
	result["describe"] = "Login Success"
	ctx.Data["json"] = result
	ctx.ServeJSON()

}
