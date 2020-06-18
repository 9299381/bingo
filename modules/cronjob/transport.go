package cronjob

import (
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
)

func NewCronJob(endpoint endpoint.Endpoint) *bingo.CommonServer {
	return bingo.NewCommonServer(
		endpoint,
		CronDecodeRequest,
		CronEncodeResponse,
	)
}
