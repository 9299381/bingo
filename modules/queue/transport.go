package queue

import (
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
)

func NewQueueHandler(endpoint endpoint.Endpoint) *bingo.CommonServer {
	return bingo.NewCommonServer(
		endpoint,
		QueueDecodeRequest,
		QueueEncodeResponse,
	)
}
