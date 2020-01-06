package controllers

import (
	"encoding/json"
	"fanqiechaodan-Accounts/models"
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
func (u *UserController) Post() {
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
	mapRet := make(map[string]string)

	if uid != "" {
		user, err := models.GetUser(uid)
		mapRet["id"] = user.Id
		mapRet["token"] = user.Profile
		if err != nil {
			u.ServerFailed(403, err.Error())
			return
		} else {
			u.ServerOk(mapRet)
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
	userInfo, err := models.Login(user.Username, user.Password)
	mapRet := make(map[string]string)

	mapRet["id"] = userInfo.Id
	mapRet["token"] = userInfo.Profile
	if err != nil {
		u.ServerFailed(403, err.Error())
		return
	} else {
		u.ServerOk(mapRet)
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
func (u *UserController) Status() {
	uid := u.GetString(":uid")
	status, err := models.GetStatus(uid)
	if err != nil {
		u.ServerFailed(400, err.Error())
		return
	} else {
		mapRet := make(map[string]int)
		mapRet["status"] = status
		u.ServerOk(mapRet)
		return
	}
}
