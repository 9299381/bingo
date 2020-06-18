package bingo

import (
	"errors"
	"fmt"
	"sync"
)

var (
	serviceMap  = make(map[string]IService)
	serviceLock sync.Mutex
)

func RegisterService(name string, service IService) {
	if _, ok := serviceMap[name]; ok {
		panic(fmt.Sprintf("service already registered: %s", name))
	}
	serviceLock.Lock()
	defer serviceLock.Unlock()
	serviceMap[name] = service
}
func Service(name string) IService {
	if service, ok := serviceMap[name]; ok {
		return service
	}
	return nil
}

//-------------------//
type IService func(ctx Context) error

func Pipe(services ...IService) *PipeService {
	var s []IService
	for _, service := range services {
		s = append(s, service)
	}
	return &PipeService{
		services: s,
	}
}

type PipeService struct {
	services []IService
}

func (s *PipeService) Len() int {
	return len(s.services)
}

func (s *PipeService) Middle(services ...IService) *PipeService {
	for _, service := range services {
		if service != nil {
			s.services = append(s.services, service)
		}
	}
	return s
}

func (s *PipeService) Call(ctx Context) error {
	if s.services[0] != nil {
		return s.services[0](ctx)
	}
	return errors.New("no service to call handle")
}

func (s *PipeService) Line(ctx Context) error {
	if len(s.services) > 0 {
		for _, service := range s.services {
			err := service(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("no service to line handle")
}
func (s *PipeService) Parallel(ctx Context) error {
	type st struct {
		Context
		err error
	}
	ch := make([]chan st, len(s.services))
	newCtx := make([]Context, len(s.services))
	for k, service := range s.services {
		ch[k] = make(chan st)
		newCtx[k], _ = WithContextValue(ctx)
		go func(cc Context, s IService, c chan st) {
			err := s(cc)
			ret := st{
				Context: cc,
				err:     err,
			}
			c <- ret
		}(newCtx[k], service, ch[k])
	}
	m := make(map[string]interface{})
	for _, c := range ch {
		res := <-c
		if res.err != nil {
			return res.err
		}
		for key, value := range res.Keys {
			m[key] = value
		}
	}
	ctx.Keys = m
	return nil
}
