package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/queue"
	"testing"
)

func TestQueue(t *testing.T) {
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
	fmt.Println(err)
}
