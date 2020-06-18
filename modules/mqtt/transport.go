package mqtt

import (
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
)

func NewMqttSubscribe(endpoint endpoint.Endpoint) *bingo.CommonServer {
	return bingo.NewCommonServer(
		endpoint,
		MQTTDecodeRequest,
		MQTTEncodeResponse,
	)
}
