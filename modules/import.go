package modules

// server,hook型
import (
	_ "github.com/9299381/bingo/modules/command"
	_ "github.com/9299381/bingo/modules/cronjob"
	_ "github.com/9299381/bingo/modules/delayq"
	_ "github.com/9299381/bingo/modules/event"
	_ "github.com/9299381/bingo/modules/grpc"
	_ "github.com/9299381/bingo/modules/http"
	_ "github.com/9299381/bingo/modules/mqtt"
	_ "github.com/9299381/bingo/modules/queue"
	_ "github.com/9299381/bingo/modules/registy"
	_ "github.com/9299381/bingo/modules/swagger"
)
