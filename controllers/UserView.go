package controllers

import "github.com/astaxie/beego"

type UserViewController struct {
	beego.Controller
}

func (c *UserViewController) Get() {

	c.TplName = "user_page.html"
}

