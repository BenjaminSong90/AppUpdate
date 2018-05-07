package controllers

import (
	"github.com/astaxie/beego"
	"AppUpdate/models"
	"AppUpdate/tools/tokenTools"
)

type UserController struct {
	beego.Controller
}

func (ctx *UserController) Post(){
	result := make(map[string]interface{})

	email := ctx.Input().Get("email")
	//valid := validation.Validation{}
	//status := valid.Email(email,email).Ok
	//if(status){
	//	result["status"] = "ERROR"
	//	result["token"] = ""
	//	result["describe"] = "Not mailbox"
	//	ctx.Data["json"] = result
	//	ctx.ServeJSON()
	//	return
	//}

	_,err := models.GetUserByEmail(email)

	if err == nil{
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "The user has already existed"
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	user := &models.User{}
	user.Email = email
	user.Active = true
	user.Password = ctx.Input().Get("password")
	user.UserName = ctx.Input().Get("username")
	_,error := models.UserAdd(user)
	if error!= nil{
		token := tokenTools.CreateToken(user.UserName)
		result["status"] = "OK"
		result["token"] = token
		result["describe"] = "The user create success"
		ctx.Data["json"] = result
		ctx.ServeJSON()
	}else{
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "User insert error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
	}

}
