package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/event"
)

func init() {
	bingo.RegisterHandler(new(EventHandler))
}

type EventHandler struct {
}

func (*EventHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.event",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(EventHandler)
		},
	}
}

func (*EventHandler) Handle(ctx bingo.Context) (interface{}, error) {

	fmt.Println("event handler")
	m := make(map[string]interface{})
	m["event"] = "event"

	payload := &bingo.Payload{
		Route: "demo.event",
		Params: map[string]interface{}{
			"payload": "payload",
			"request": "test",
		},
	}
	err := event.Fire(payload)
	if err != nil {
		return nil, err
	}

	return m, nil
}

var (
	_ bingo.IHandler = (*EventHandler)(nil)
)
