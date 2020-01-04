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

	beego.Router("/v1/user", &controllers.UserController{})
	beego.Router("/v1/user/:uid", &controllers.UserController{}, "get:GetUser")
	beego.Router("/v1/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/v1/user/logout", &controllers.UserController{}, "get:Logout")

}
