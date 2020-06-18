package handler

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/consul"
)

func init() {
	bingo.RegisterHandler(new(ConsulHandler))
}

type ConsulHandler struct{}

func (*ConsulHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.consul",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(ConsulHandler)
		},
	}
}

func (*ConsulHandler) Handle(ctx bingo.Context) (interface{}, error) {
	entity, _ := consul.GetService("bingo")
	ctx.Log.Info(entity.Service.Service)
	ctx.Log.Info(entity.Service.Address)
	ctx.Log.Info(entity.Service.Port)
	if tag := entity.Service.Tags[0]; tag != "" {
		return tag, nil
	}
	return nil, nil
}

var (
	_ bingo.IHandler = (*ConsulHandler)(nil)
)
