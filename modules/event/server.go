package event

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/id"
	"github.com/go-kit/kit/endpoint"
	"runtime"
	"sync"
	"time"
)

func init() {
	bingo.RegisterModule(new(Server))
}

type Server struct {
	bingo.Context
	Concurrency int
	After       <-chan time.Time

	eventPool sync.Pool
	handlers  map[string]endpoint.Endpoint
	eventChan chan *bingo.Payload
	mutex     sync.RWMutex
}

func (s *Server) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "event",
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
	s.Concurrency = config.EnvInt("event.concurrency", 1)
	after := config.EnvInt("event.after", 1)
	s.After = time.After(time.Duration(after) * time.Second)
	//
	s.handlers = map[string]endpoint.Endpoint{}
	s.eventChan = make(chan *bingo.Payload, runtime.NumCPU())
}

func (s *Server) Route(name string, endpoint endpoint.Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.handlers[name] = endpoint
}

func (s *Server) GetRoute(name string) endpoint.Endpoint {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if ret, ok := s.handlers[name]; ok {
		return ret
	}
	return nil
}
func (s *Server) AddEvent(payload *bingo.Payload) error {
	event := s.newEvent(payload)
	go func(e *bingo.Payload) {
		s.eventChan <- e
	}(event)
	return nil
}

func (s *Server) Start(id string) error {
	s.Log.Infof("%s start at Concurrency %d ", id, s.Concurrency)
	errChan := make(chan error)
	for i := 0; i < s.Concurrency; i++ {
		go s.handleEventReceive(errChan)
	}
	err := <-errChan
	if err != nil {
		s.Log.Info(err)
		return err
	}
	return nil

}

func (s *Server) handleEventReceive(errChan chan error) {
	for {
		select {
		case event := <-s.eventChan:
			handler := s.GetRoute(event.Route)
			if handler != nil {
				ctx := context.Background()
				request := &bingo.Request{
					Id:   id.New(),
					Data: event.Params,
				}
				resp, err := handler(ctx, request)
				if err != nil {
					s.eventPool.Put(event)
					s.Log.Info("event error:", err)
					//errChan <- err // 退出协程了
				} else {
					s.Log.Info("event response:", resp)
				}
			}
			s.eventPool.Put(event)
		case <-s.After:
			s.Log.Info("event wait ......")
		}
	}
}

func (s *Server) Stop(id string) {
	s.Log.Infof("%s stop now", id)
}

func (s *Server) newEvent(payload *bingo.Payload) *bingo.Payload {
	e := s.eventPool.Get()
	if e == nil {
		return payload
	} else {
		ret := e.(*bingo.Payload)
		(*ret).Route = payload.Route
		(*ret).Params = payload.Params
		return ret
	}
}

var (
	_ bingo.IModule       = (*Server)(nil)
	_ bingo.IModuleServer = (*Server)(nil)
)
