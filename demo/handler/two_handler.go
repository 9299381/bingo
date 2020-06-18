package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/demo/service"
)

func init() {
	bingo.RegisterHandler(new(TwoHandler))
}

type TwoHandler struct {
}

func (*TwoHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.two",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(TwoHandler)
		},
	}
}

func (*TwoHandler) Handle(ctx bingo.Context) (interface{}, error) {
	fmt.Println("two handler")
	m := make(map[string]interface{})
	m["handler"] = "two_handler"
	fmt.Println(ctx.Request)
	err := bingo.Pipe().Middle(
		service.OneService,
		service.TwoService,
	).Line(ctx)
	if err != nil {
		return nil, err
	}
	m["aaa"] = ctx.GetString("aaa")
	m["bbb"] = ctx.GetString("bbbbbb")
	return m, nil
}

var (
	_ bingo.IHandler = (*TwoHandler)(nil)
)
