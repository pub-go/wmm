package model

type Response struct {
	Code     int64  `json:"code"`
	Messsage string `json:"messsage"`
	Data     any    `json:"data"`
}
