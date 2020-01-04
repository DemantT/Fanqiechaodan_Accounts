package controllers

import (
	"fanqiechaodan-Accounts/models"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) ServerOk(user models.User) {
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *ErrorController) ServerFailed(code int, errMsg string) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = errMsg
	c.ServeJSON()
}
