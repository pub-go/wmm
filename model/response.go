package model

type Response[T any] struct {
	Code     int64  `json:"code"`
	Messsage string `json:"messsage"`
	Data     T      `json:"data"`
}
