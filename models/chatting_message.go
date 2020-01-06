package models

type ChattingMessage struct {
	UserId int64 `json:"user_id"`
	UserYype string `json:"user_type"`
	Message string `json:"meesage"`
	Code string `json:"code"`
}


func (m ChattingMessage) History(userId int64) []ChattingMessage {
	filename := "chatting_message_" + string(userId)
	err, chattingMessages := ReadChattingMessages(filename)
	if err != nil {
		return nil
	}
	return chattingMessages
}

func (m ChattingMessage) CreateChattingMessage(message ChattingMessage) []ChattingMessage {
	filename := "chatting_message_" + string(message.UserId)
	err, chattingMessages := WriteChattingMessage(filename, message)
	if err != nil {
		return nil
	}
	return chattingMessages
}

