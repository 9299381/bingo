package grpc

import (
	"context"
	"encoding/json"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/grpc/protobuf"
)

// TCP请求数据解码函数
func GprcDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*protobuf.Request)
	data := make(map[string]interface{})
	_ = json.Unmarshal([]byte(request.Param), &data)
	return &bingo.Request{
		Id:   request.Id,
		Data: data,
	}, nil
}

// TCP返回数据编码函数
func GprcEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	response := rsp.(*bingo.Response)
	data, _ := json.Marshal(response.Data)
	resp := &protobuf.Response{
		Code: response.Code,
		Data: string(data),
		Msg:  response.Message,
	}
	return resp, nil
}
