package controllers

import (
	"encoding/json"
	"fanqiechaodan-Accounts/models"
	"fmt"
)

// Operations about Users
type UserController struct {
	ErrorController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Create() {
	var user models.User
	mapRet := make(map[string]string)
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	userInfo, err := models.AddUser(user)
	mapRet["id"] = userInfo.Id
	mapRet["token"] = userInfo.Profile
	if err != nil {
		u.ServerFailed(500, err.Error())
	} else {
		u.ServerOk(mapRet)
	}
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) GetUser() {
	uid := u.GetString(":uid")
	//mapRet := make(map[string]string)

	if uid != "" {
		user, err := models.GetUser(uid)
		resp := new(models.Resp)

		resp.Meta = models.Meta{
			Code:    20000,
			Type:    "",
			Message: "",
		}

		mapRet := make(map[string]interface{})

		mapRet["id"] = user.Id
		mapRet["token"] = user.Profile
		resp.Data = mapRet
		if err != nil {
			u.ServerFailed(403, err.Error())
			return
		} else {
			u.ServerOk(resp)
			return
		}
	}
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	userInfo, err, bIsnew := models.Login(user.Username, user.Password)
	resp := new(models.Resp)

	resp.Meta = models.Meta{
		Code:    20000,
		Type:    "",
		Message: "",
	}

	mapRet := make(map[string]interface{})

	mapRet["id"] = userInfo.Id
	mapRet["token"] = userInfo.Profile
	resp.Data = mapRet
	if err != nil {
		u.ServerFailed(403, err.Error())
		return
	} else {
		if !bIsnew {
			models.WriteUser("/Users/liuchuan/myowncode/testcode/test.json", *userInfo)
		}
		u.ServerOk(resp)
		return
	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.ServerOk("logout success")
}

// @Title status
// @Description status in user session
// @Success 200 {status:number}
// @router /status/:uid [get]
func (u *UserController) GetStatus() {
	uid := u.GetString(":uid")
	status, err := models.GetStatus(uid)
	if err != nil {
		fmt.Println("status err is ", err)
		u.ServerFailed(400, err.Error())
		return
	}
	resp := new(models.Resp)

	resp.Meta = models.Meta{
		Code:    20000,
		Type:    "",
		Message: "",
	}

	mapRet := make(map[string]interface{})

	mapRet["status"] = status
	resp.Data = mapRet

	if err != nil {
		u.ServerFailed(400, err.Error())
		return
	} else {
		u.ServerOk(resp)
		return
	}
}

// @Title sendmail
// @Description send mail
// @Success 200
// @router /mail/send [post]
func (u *UserController) SendMail() {
	var mail models.Mail
	uid := u.GetString(":uid")
	json.Unmarshal(u.Ctx.Input.RequestBody, &mail)
	err := models.SendMail(uid, mail)
	if err != nil {
		fmt.Println("send mail err is ", err)
		u.ServerFailed(400, err.Error())
		return
	}
	resp := new(models.Resp)
	resp.Meta = models.Meta{
		Code:    20000,
		Type:    "",
		Message: "",
	}
	resp.Data = ""

	if err != nil {
		u.ServerFailed(400, err.Error())
		return
	} else {
		u.ServerOk(resp)
		return
	}
}
