package controllers

import (
	"encoding/json"
	"fanqiechaodan-Accounts/models"
)

// Operations about Users
type BotChattingController struct {
	ErrorController
}

func (u *BotChattingController) Send() {
	var botChatting models.BotChatting
	json.Unmarshal(u.Ctx.Input.RequestBody, &botChatting)
	reply, err := models.Chat(botChatting.Message)
	if err != nil {

	}
	resp := new(models.Resp)
	resp.Meta = models.Meta{
		Code:    20000,
		Type:    "",
		Message: "",
	}

	mapRet := make(map[string]interface{})

	mapRet["reply"] = reply
	resp.Data = mapRet
	if err != nil {
		u.ServerFailed(403, err.Error())
		return
	} else {
		u.ServerOk(resp)
		return
	}
}
