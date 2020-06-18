package handler

import "github.com/9299381/bingo"

func init() {
	bingo.RegisterHandler(new(ValidHandler))
}

type ValidHandler struct{}

func (*ValidHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.valid",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(ValidHandler)
		},
	}
}

func (*ValidHandler) Handle(ctx bingo.Context) (interface{}, error) {

	return nil, nil
}

var (
	_ bingo.IHandler = (*ValidHandler)(nil)
)
