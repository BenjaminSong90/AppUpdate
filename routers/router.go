package routers

import (
	"AppUpdate/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.UserController{})
    beego.Router("/login",&controllers.LoginController{})
}
