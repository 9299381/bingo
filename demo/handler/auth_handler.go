package handler

import (
	"fmt"
	"github.com/9299381/bingo"
)

func init() {
	bingo.RegisterHandler(new(AuthHandler))
}

type AuthHandler struct {
}

func (handler *AuthHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.auth",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(AuthHandler)
		},
	}
}

func (*AuthHandler) Handle(ctx bingo.Context) (interface{}, error) {
	fmt.Println("auth handler")
	fmt.Println(ctx.Request.Data["claim"])
	return nil, nil
}

var (
	_ bingo.IHandler = (*AuthHandler)(nil)
)
