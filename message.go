package bingo

import (
	"strings"
)

type Request struct {
	Id   string `json:"request_id"`
	Data map[string]interface{}
}

func (req *Request) Claim(key string) (s string) {
	if m, ok := req.Data["claim"].(map[string]interface{}); ok {
		if val, has := m[key]; has {
			s, _ = val.(string)
		}
	}
	return
}

func (req *Request) Set(key string, value interface{}) {
	req.Data[key] = value
}

func (req *Request) Get(key string) interface{} {
	value, has := req.Data[key]
	if has {
		return value
	} else {
		return nil
	}
}

func (req *Request) GetString(key string) (s string) {
	if val := req.Get(key); val != nil {
		s, _ = val.(string)
	}
	return
}
func (req *Request) GetInt(key string) (s int) {
	if val := req.Get(key); val != nil {
		s, _ = val.(int)
	}
	return
}
func (req *Request) GetFloat(key string) (s float64) {
	if val := req.Get(key); val != nil {
		s, _ = val.(float64)
	}
	return
}

type Response struct {
	Ret     int         `json:"ret"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func MakeResponse(data interface{}, err error) *Response {
	if err != nil {
		return Failed(err)
	} else {
		return Success(data)
	}
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    "0000",
		Data:    data,
		Ret:     200,
		Message: "请求成功",
	}
}
func Failed(err error) *Response {
	errMap := strings.Split(err.Error(), "::")
	if len(errMap) == 2 {
		return &Response{
			Code:    errMap[0],
			Data:    make(map[string]interface{}),
			Ret:     200,
			Message: errMap[1],
		}
	} else {
		return &Response{
			Code:    "9999",
			Data:    make(map[string]interface{}),
			Ret:     200,
			Message: err.Error(),
		}
	}
}

type Payload struct {
	Route  string                 `json:"route"`
	Params map[string]interface{} `json:"params"`
}

type Job struct {
	Queue  string                 `json:"queue"`
	Route  string                 `json:"route"`
	Params map[string]interface{} `json:"params"`
}
type Code struct {
	Key   string
	Value string
	Msg   string
}
