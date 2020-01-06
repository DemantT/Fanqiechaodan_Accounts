// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fanqiechaodan-Accounts/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/create", &controllers.UserController{}, "post:Create"),
			beego.NSRouter("/:uid", &controllers.UserController{}, "get:GetUser"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "get:Logout"),
			beego.NSRouter("/status/:uid", &controllers.UserController{}, "get:GetStatus"),
		),
		beego.NSNamespace("/messages"), //beego.NSRouter("/history/:id", &controllers.UserController{}, "get:History"),
		//beego.NSRouter("/send", &controllers.UserController{}, "post:Send"),

	)
	beego.AddNamespace(ns)
}
