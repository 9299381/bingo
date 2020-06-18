package provider

import (
	"github.com/9299381/bingo"
	_ "github.com/9299381/bingo/demo/handler"
)

func init() {

	registerMiddleware()
	registerProvider()

}

func registerMiddleware() {
	bingo.RegisterMiddleware("auth",
		bingo.Middleware(
			bingo.AuthMiddleware(),
		),
	)

}

func registerProvider() {
	bingo.Provider(&HttpProvider{})
	bingo.Provider(&GrpcProvider{})
	bingo.Provider(&CronProvider{})
	bingo.Provider(&EventProvider{})
	bingo.Provider(&QueueProvider{})
	bingo.Provider(&MQTTProvider{})
	bingo.Provider(&CmdProvider{})
}
