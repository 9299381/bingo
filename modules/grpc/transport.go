package grpc

import (
	"github.com/go-kit/kit/endpoint"
	GrpcTransport "github.com/go-kit/kit/transport/grpc"
)

func NewGRPC(endpoint endpoint.Endpoint) *GrpcTransport.Server {
	return GrpcTransport.NewServer(
		endpoint,
		GprcDecodeRequest,
		GprcEncodeResponse,
	)
}
