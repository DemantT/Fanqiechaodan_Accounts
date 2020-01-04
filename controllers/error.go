package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) ServerOk(retMsg interface{}) {
	c.Data["json"] = retMsg
	c.ServeJSON()
}

func (c *ErrorController) ServerFailed(code int, errMsg string) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Data["json"] = errMsg
	c.ServeJSON()
}
