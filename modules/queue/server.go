package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/redis"
	"github.com/go-kit/kit/endpoint"
	redigo "github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

func init() {
	bingo.RegisterModule(new(Server))
}

type Options struct {
	Prefix      string
	Listen      []string
	Interval    time.Duration
	Concurrency int
	UseNumber   bool
}

type Server struct {
	opts *Options
	bingo.Context
	handlers map[string]bingo.ICommonHandler
	mutex    sync.RWMutex
}

func (s *Server) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "queue",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0.0",
		New: func() bingo.IModule {
			server := new(Server)
			server.ConfigModule()
			return server
		},
	}
}

func (s *Server) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	s.handlers = map[string]bingo.ICommonHandler{}
	s.opts = &Options{
		Prefix:      config.EnvString("queue.prefix", "bingo"),
		Listen:      config.EnvStringSlice("queue.listen", []string{}),
		Interval:    time.Duration(config.EnvInt("queue.interval", 1)) * time.Second,
		Concurrency: config.EnvInt("queue.concurrency", 1),
		UseNumber:   true,
	}

}
func (s *Server) Route(name string, endpoint endpoint.Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	handler := NewQueueHandler(endpoint)
	s.handlers[name] = handler
}

func (s *Server) getHandler(name string) bingo.ICommonHandler {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.handlers[name]
}

func (s *Server) Start(id string) error {
	if redis.Pool() == nil {
		return errors.New("redis connection error")
	}
	s.Log.Infof("%s start ...", id)
	errChan := make(chan error)
	jobs := s.poll(errChan)
	for id := 0; id < s.opts.Concurrency; id++ {
		s.work(id, jobs, errChan)
	}
	return <-errChan
}

func (s *Server) work(id int, jobs <-chan *bingo.Job, errChan chan error) {
	go func() {
		for job := range jobs {
			if handler := s.getHandler(job.Route); handler != nil {
				ctx := context.Background()
				response, err := handler.ServeHandle(ctx, job)
				if err != nil {
					errChan <- err
					return
				}
				s.Log.Debugf("Concurrency_Id:%d ,Job Response:%v", id, response)
			} else {
				errorLog := fmt.Sprintf(
					"No worker for %s in queue %s with args %v",
					job.Route,
					job.Queue,
					job.Params)
				s.Log.Error(errorLog)
				errChan <- errors.New(errorLog)
				return
			}
		}
	}()
}

func (s *Server) poll(errChan chan error) <-chan *bingo.Job {
	jobs := make(chan *bingo.Job)
	go func() {
		conn := redis.Pool().Get()
		defer conn.Close()
		for {
			select {
			default:
				job, err := s.getJob(conn)
				if err != nil {
					errorLog := fmt.Sprintf(
						"Error on %v getting job from: %v",
						s.opts.Listen,
						err)
					s.Log.Error(errorLog)
					errChan <- errors.New(errorLog)
					return
				}
				if job != nil {
					jobs <- job
				} else {
					s.Log.Debugf("Waiting for %v", s.opts.Listen)
					timeout := time.After(s.opts.Interval)
					select {
					case <-timeout:
					}
				}
			}
		}
	}()

	return jobs
}
func (s *Server) getJob(conn redigo.Conn) (*bingo.Job, error) {
	for _, queue := range s.opts.Listen {
		arg := fmt.Sprintf("%s_queue:%s", s.opts.Prefix, queue)
		reply, err := conn.Do("LPOP", arg)
		if err != nil {
			return nil, err
		}
		if reply != nil {
			s.Log.Infof("Found job on %s", queue)
			job := &bingo.Job{Queue: queue}
			decoder := json.NewDecoder(bytes.NewReader(reply.([]byte)))
			if s.opts.UseNumber {
				decoder.UseNumber()
			}
			if err := decoder.Decode(&job); err != nil {
				return nil, err
			}
			return job, nil
		}
	}
	return nil, nil
}

func (s *Server) Stop(id string) {
	s.Log.Infof("%s stop ...", id)
	//如果jobs 不为空, 发回redis ?
}

var (
	_ bingo.IModule       = (*Server)(nil)
	_ bingo.IModuleServer = (*Server)(nil)
)
