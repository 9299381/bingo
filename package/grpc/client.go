package grpc

import (
	"context"
	"encoding/json"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/grpc/protobuf"
	"github.com/9299381/bingo/package/id"
	"github.com/9299381/bingo/package/logger"
	"google.golang.org/grpc"
)

func Call(host, service string, params map[string]interface{}) *bingo.Response {
	resp, err := newClient(host, service, params)
	if err != nil {
		return bingo.Failed(err)
	}
	m := make(map[string]interface{})
	m["call_method"] = "grpc"
	err = json.Unmarshal([]byte(resp.GetData()), &m)
	if err != nil {
		return bingo.Failed(err)
	}
	return &bingo.Response{
		Ret:     200,
		Code:    resp.Code,
		Data:    m,
		Message: resp.Msg,
	}
}
func newClient(serviceAddress string, service string, params map[string]interface{}) (*protobuf.Response, error) {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		logger.GetInstance().Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	jsonParam, _ := json.Marshal(params)
	in := &protobuf.Request{
		Id:    id.New(),
		Param: string(jsonParam),
	}

	out := new(protobuf.Response)

	method := "/protobuf." + service + "/Handle"
	err = conn.Invoke(context.Background(), method, in, out)
	return out, err
}
