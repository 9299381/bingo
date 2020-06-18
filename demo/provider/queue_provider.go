package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/queue"
)

type QueueProvider struct {
}

func (q *QueueProvider) Boot() {
}

func (q *QueueProvider) Register() {
	bingo.Bind("queue", func(module bingo.IModule) error {
		mod := module.(*queue.Server)
		mod.Route("demo.queue", bingo.Handler("demo.two"))
		return nil
	})
}
