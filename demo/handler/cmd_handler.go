package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
)

func init() {
	bingo.RegisterHandler(new(CmdHandler))
}

type CmdHandler struct {
}

func (*CmdHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.cmd",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(CmdHandler)
		},
	}
}

func (*CmdHandler) Handle(ctx bingo.Context) (interface{}, error) {
	fmt.Println("cmd handler")
	fmt.Println(ctx.Request.GetString("hello"))
	fmt.Println(config.EnvString("server.http_port", "123"))
	return *ctx.Request, nil
}

var (
	_ bingo.IHandler = (*CmdHandler)(nil)
)
