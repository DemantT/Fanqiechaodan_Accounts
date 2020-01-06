package controllers

import (
	"encoding/json"
	"fanqiechaodan-Accounts/models"
	"fmt"
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
	fmt.Println(reply)
}
