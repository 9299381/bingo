package queue

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
)

func QueueDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	job := req.(*bingo.Job)
	data := make(map[string]interface{})
	data["queue"] = job.Queue
	data["route"] = job.Route
	data["params"] = job.Params
	return &bingo.Request{
		Id:   id.New(),
		Data: data,
	}, nil
}

func QueueEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
