package command

import (
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
)

func NewCommandHandler(endpoint endpoint.Endpoint) *bingo.CommonServer {
	return bingo.NewCommonServer(
		endpoint,
		CommandDecodeRequest,
		CommandEncodeResponse,
	)
}
