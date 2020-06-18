package cronjob

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
)

func CronDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	request := req.(map[string]interface{})
	return &bingo.Request{
		Id:   id.New(),
		Data: request,
	}, nil
}

func CronEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
