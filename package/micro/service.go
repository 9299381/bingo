package micro

import (
	"context"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/consul"
	"github.com/9299381/bingo/package/grpc"
	"github.com/9299381/bingo/package/http"
	"github.com/9299381/bingo/package/id"
)

// resp *bingo.Response = micro.Service(name).api(api).Request(params).Run()
// micro -> service,  service ->route
func Service(service string) *microService {
	return &microService{
		service: service,
		params:  make(map[string]interface{}),
	}
}

type microService struct {
	service string
	api     string
	params  map[string]interface{}
}

func (s *microService) Api(api string) *microService {
	s.api = api
	return s
}

func (s *microService) Request(params map[string]interface{}) *microService {
	s.params = params
	return s
}

func (s *microService) Run() *bingo.Response {

	microType := config.EnvString("micro_type", "local")
	if microType == "local" {
		return s.localRun()
	} else if microType == "consul" {
		return s.consulRun()
	}
	return nil
}

func (s *microService) localRun() *bingo.Response {
	req := &bingo.Request{
		Id:   id.New(),
		Data: s.params,
	}
	if handler := bingo.Handler(s.api); handler != nil {
		ret, err := handler(context.Background(), req)
		if err != nil {
			return bingo.Failed(err)
		}
		return ret.(*bingo.Response)
	} else {
		return bingo.Failed(fmt.Errorf("no handler for %s", s.api))
	}

}

func (s *microService) consulRun() *bingo.Response {
	entity, err := consul.GetService(s.service)
	if err != nil {
		return bingo.Failed(err)
	}
	tag := entity.Service.Tags[0]
	host := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
	//
	if tag == "http" {
		return http.Post(host, s.api, s.params)
	} else if tag == "grpc" {
		return grpc.Call(host, s.api, s.params)
	}
	return bingo.Failed(fmt.Errorf("no handler for %s", s.api))

}
