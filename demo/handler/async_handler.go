package handler

import (
	"github.com/9299381/bingo"
	"time"
)

func init() {
	bingo.RegisterHandler(new(AsyncHandler))
}

type AsyncHandler struct{}

func (AsyncHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.async_handler",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(AsyncHandler)
		},
	}
}

func (AsyncHandler) Handle(ctx bingo.Context) (interface{}, error) {
	go func() {
		for i := 0; i < 10; i++ {
			ctx.Log.Info("async handle ....")
			time.Sleep(1 * time.Second)
		}
	}()
	return "ok", nil
}
