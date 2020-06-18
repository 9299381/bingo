package test

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/queue"
	"testing"
)

func TestDelayQueue(t *testing.T) {

	job := &bingo.Job{
		Queue: "demo1",
		Route: "demo.queue",
		Params: map[string]interface{}{
			"hello": "world",
			"num":   1,
		},
	}
	_ = queue.Delay(600, job)
}
func TestUnDelayQueue(t *testing.T) {
	job := &bingo.Job{
		Queue: "demo1",
		Route: "demo.queue",
		Params: map[string]interface{}{
			"hello": "world",
			"num":   "KbsE",
		},
	}
	_ = queue.UnDelay(job)
}
