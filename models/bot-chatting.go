package models

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type BotChatting struct {
	Message string `json:"message"`
}

type BaiduRequest struct {
	LogId       string      `json:"log_id"`
	Version     string      `json:"version"`
	ServiceId   string      `json:"service_id"`
	SessionId   string      `json:"session_id"`
	Request     Request     `json:"request"`
	DialogState DialogState `json:"dialog_state"`
}

type Request struct {
	Query  string `json:"query"`
	UserId string `json:"user_id"`
}

type DialogState struct {
	Contexts Contexts `json:"contexts"`
}

type Contexts struct {
	SYS_REMEMBERED_SKILLS []string `json:"SYS_REMEMBERED_SKILLS"`
}

type BaiduResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	ResponseList []ResponseItem `json:"response_list"`
}

type ResponseItem struct {
	ActionList []ActionItem `json:"action_list"`
}

type ActionItem struct {
	Say string `json:"say"`
}

func Chat(message string) (reply string, err error) {
	token := "24.54ae1b54c5927fac0569527519f01ed0.2592000.1580622944.282335-18167418"
	requestStruct := BaiduRequest{
		LogId:     "d88a4c70-2dee-11ea-abcc-6f8ad35862a1",
		Version:   "2.0",
		ServiceId: "S25897",
		SessionId: "service-session-id-1578031421469-4ab17bffc80b4cc3ac31c41c413f7a87",
		Request: Request{
			Query:  message,
			UserId: "88888",
		},
		DialogState: DialogState{
			Contexts: Contexts{
				SYS_REMEMBERED_SKILLS: []string{"1012568"},
			},
		},
	}

	respone := Post("https://aip.baidubce.com/rpc/2.0/unit/service/chat?access_token="+token, requestStruct, "application/x-www-form-urlencoded")

	var baiduResponse BaiduResponse
	err = json.Unmarshal(respone, &baiduResponse)
	if err != nil {
	}

	return baiduResponse.Result.ResponseList[0].ActionList[0].Say, nil
}

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	response = result.String()
	return
}

//发送POST请求
//url:请求地址		data:POST请求提交的数据		contentType:请求体格式，如：application/json
//content:请求返回的内容
func Post(url string, data interface{}, contentType string) (response []byte) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result
}
