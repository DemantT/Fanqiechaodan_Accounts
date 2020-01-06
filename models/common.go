package models

type Meta struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Resp struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
