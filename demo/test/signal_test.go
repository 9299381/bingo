package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"os"
	"testing"
	"time"
)

var jobs = make(chan *bingo.Job, 10)
var b = make(chan bool)

func TestSignal(t *testing.T) {
	job := &bingo.Job{
		Queue:  "demo",
		Route:  "demo",
		Params: nil,
	}
	jobs <- job
	jobs <- job
	jobs <- job
	close(jobs)
	for i := 0; i < cap(jobs)+1; i++ {
		v, ok := <-jobs
		fmt.Println(v, ok)
	}

	//quit:= make(chan error)
	//bingo.Signals(ExitFunc)
	//
	//go getJobs(quit)
	//for job:= range jobs{
	//	fmt.Println(job.Route)
	//	//time.Sleep(time.Duration(2)*time.Second)
	//}
	//
	//ExitFunc(<-quit)
}

func getJobs(quit chan error) {
	for {
		job := &bingo.Job{
			Queue:  "demo",
			Route:  "demo",
			Params: nil,
		}
		select {
		case jobs <- job:
		case <-b:
			break
		case <-time.After(time.Duration(1) * time.Second):
		}
	}
}
func ExitFunc(err error) {
	close(b)
	fmt.Println("开始退出...")
	//for job := range jobs{
	//	fmt.Println(job)
	//}
	//close(jobs)
	fmt.Println("结束退出...")
	os.Exit(0)
}
