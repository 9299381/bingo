package handler

import (
	"fmt"
	"github.com/9299381/bingo"
)

func init() {
	bingo.RegisterHandler(new(CronHandler))
}

type CronHandler struct {
}

func (*CronHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.cron",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(CronHandler)
		},
	}
}

func (*CronHandler) Handle(ctx bingo.Context) (interface{}, error) {

	fmt.Println("run cron job here!!!")
	return nil, nil
}

var (
	_ bingo.IHandler = (*CronHandler)(nil)
)
