package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/queue"
)

func init() {
	bingo.RegisterHandler(new(QueueHandler))
}

type QueueHandler struct {
}

func (*QueueHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.queue",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(QueueHandler)
		},
	}
}

func (*QueueHandler) Handle(ctx bingo.Context) (interface{}, error) {
	fmt.Println("queue_handler")
	m := make(map[string]interface{})
	m["handler"] = "queue_handler"
	//这里的demo1 是在配置文件中设置用于 队列侦听的
	// queue.one 是侦听服务器绑定的路由
	job := &bingo.Job{
		Queue:  "demo1",
		Route:  "demo.queue",
		Params: m,
	}
	err := queue.Fire(job)
	if err != nil {
		return nil, err
	}
	return m, nil
}

var (
	_ bingo.IHandler = (*QueueHandler)(nil)
)
