package mqtt

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MQTTDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	request := req.(mqtt.Message)
	data := make(map[string]interface{})
	data["topic"] = request.Topic()
	var payload map[string]interface{}
	err := json.Unmarshal(request.Payload(), &payload)
	if err != nil {
		return nil, errors.New("mqtt payload json error")
	}
	data["payload"] = payload
	requestId, ok := payload["request_id"].(string)
	if ok == false {
		requestId = id.New()
	}
	return &bingo.Request{
		Id:   requestId,
		Data: data,
	}, nil
}

func MQTTEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
