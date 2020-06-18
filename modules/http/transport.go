package http

import (
	"github.com/go-kit/kit/endpoint"
	HttpTransport "github.com/go-kit/kit/transport/http"
)

func NewHTTP(endpoint endpoint.Endpoint) *HttpTransport.Server {
	return HttpTransport.NewServer(
		endpoint,
		HttpFormDecodeRequest,
		HttpEncodeResponse,
	)
}

func NewWeChatNotify(endpoint endpoint.Endpoint) *HttpTransport.Server {
	return HttpTransport.NewServer(
		endpoint,
		WeChatNotifyDecodeRequest,
		WeChatNotifyEncodeResponse,
	)
}
func NewUpload(endpoint endpoint.Endpoint) *HttpTransport.Server {
	return HttpTransport.NewServer(
		endpoint,
		UploadDecodeRequest,
		UploadEncodeResponse,
	)
}
