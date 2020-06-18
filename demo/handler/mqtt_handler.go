package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"time"
)

func init() {
	bingo.RegisterHandler(new(MQTTHandler))
}

type MQTTHandler struct {
}

func (*MQTTHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.mqtt",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(MQTTHandler)
		},
	}
}

func (*MQTTHandler) Handle(ctx bingo.Context) (interface{}, error) {
	payload := ctx.Request.Data["payload"].(map[string]interface{})
	if payload["connected_at"] != nil {
		connect := int64(payload["connected_at"].(float64))
		conn := time.Unix(connect, 0).Format(bingo.YmdHis)
		fmt.Println(payload["clientid"].(string), "connect_at", conn)
	}

	if payload["disconnected_at"] != nil {
		connect := int64(payload["disconnected_at"].(float64))
		conn := time.Unix(connect, 0).Format(bingo.YmdHis)
		fmt.Println(payload["clientid"].(string), "disconnected_at", conn)
	}
	return nil, nil
}

var (
	_ bingo.IHandler = (*MQTTHandler)(nil)
)
