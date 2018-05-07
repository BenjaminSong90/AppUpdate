package main

import (
	_ "AppUpdate/routers"
	"github.com/astaxie/beego"
	_ "AppUpdate/models"
)

func main() {
	beego.Run()
}

