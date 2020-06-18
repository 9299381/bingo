package command

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
)

func CommandDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	data := make(map[string]interface{})
	if req.(string) != "" {
		err := json.Unmarshal([]byte(req.(string)), &data)
		if err != nil {
			return nil, errors.New("args参数json解析错误")
		}
	}
	return &bingo.Request{
		Id:   id.New(),
		Data: data,
	}, nil
}

func CommandEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
