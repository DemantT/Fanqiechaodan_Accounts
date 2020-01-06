package controllers

import (
	"fanqiechaodan-Accounts/models"
	"encoding/json"
)

type ChattingMessagesController struct {
	ErrorController
}

func (c *ChattingMessagesController) History() {
	var message models.ChattingMessage

	user_id := c.GetString(":id")
	filename := "chatting_message_" + string(user_id)
	models.ReadChattingMessages(filename)
	c.Data["json"] = message
	c.ServeJSON()
}

func (c *ChattingMessagesController) Send() {
	var chattingMessage models.ChattingMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &chattingMessage)
	chattingMessage.CreateChattingMessage(chattingMessage)
	
	// var robot models.RobotGirl
	// robot.reply(chattingMessage.code)

	c.ServeJSON()
}

func (c *ChattingMessagesController) ReplyOption() {
	var robot models.RobotGirl
	replyOption := robot.ReplyOption()
	c.Data["json"] = replyOption
	c.ServeJSON()
}
