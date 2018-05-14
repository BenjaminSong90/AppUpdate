package controllers

import (
	"github.com/astaxie/beego"
	"AppUpdate/models"
	"AppUpdate/tools/tokenTools"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	beego.Controller
}

type UserResult struct {
	UserName 	string
	UserEmail	string
	IsActive	bool
}

func (ctx *UserController) Post() {
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

	_, err := models.GetUserByEmail(email)


	if err == nil {
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
	spassword, _ := bcrypt.GenerateFromPassword([]byte(ctx.Input().Get("password")), bcrypt.DefaultCost)
	user.Password = string(spassword)
	user.UserName = ctx.Input().Get("username")
	_, error := models.UserAdd(user)

	beego.Debug(error == nil)
	if error == nil {
		token := tokenTools.CreateToken(user.Email)
		result["status"] = "OK"
		result["token"] = token
		result["describe"] = "The user create success"
		ctx.Data["json"] = result
		ctx.ServeJSON()
	} else {
		result["status"] = "ERROR"
		result["token"] = ""
		result["describe"] = "User insert error"
		ctx.Data["json"] = result
		ctx.ServeJSON()
	}
}

func (ctx *UserController) Get() {
	result := make(map[string]interface{})

	token := ctx.Ctx.Input.Header("token")
	beego.Debug("token:  "+token)

	if token == ""{
		result["status"] = "ERROR"
		result["describe"] = "token is nil"
		result["user_info"] = nil
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	userEmail, err := tokenTools.CheckToken(token)

	beego.Debug("userEmail:  "+userEmail)
	if err != nil {
		result["status"] = "ERROR"
		result["describe"] = "token is error"
		result["user_info"] = nil
		ctx.Data["json"] = result
		ctx.ServeJSON()
		return
	}

	user, err := models.GetUserByEmail(userEmail)
	resultUser := &UserResult{}
	resultUser.IsActive = user.Active
	resultUser.UserName = user.UserName
	resultUser.UserEmail = user.Email

	result["status"] = "SUCCESS"
	result["describe"] = "get user info success"
	result["user_info"] = resultUser

	ctx.Data["json"] = resultUser
	ctx.ServeJSON()
}


