package main

import (
	"fanqiechaodan-Accounts/controllers"
	_ "fanqiechaodan-Accounts/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
