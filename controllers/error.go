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

func (c *ErrorController) ServerFailed(err error) {
	c.Data["json"] = err
	c.ServeJSON()
}
