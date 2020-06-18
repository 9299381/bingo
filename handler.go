package bingo

import (
	"context"
	"fmt"
	"github.com/9299381/bingo/package/logger"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

var (
	handlerInfoMap   = make(map[string]*HandlerInfo)
	handlerInfoMutex sync.RWMutex
)

//------------------//
type IHandler interface {
	Info() *HandlerInfo
	Handle(ctx Context) (interface{}, error)
}

//-----------------//
type IValid interface {
	Valid(ctx Context) error
}

//------------------//
type IMock interface {
	Mock(ctx Context) interface{}
}

type HandlerInfo struct {
	ID      string
	Version string //版本,根据版本高低,替换掉注册,todo
	New     func() IHandler
}

func RegisterHandler(h IHandler) {
	info := h.Info()
	if info.ID == "" {
		panic("handler ID missing")
	}
	if info.New == nil {
		panic("missing info.New")
	}
	if _, ok := handlerInfoMap[info.ID]; ok {
		panic(fmt.Sprintf("handler already registered: %s", info.ID))
	}
	handlerInfoMutex.Lock()
	defer handlerInfoMutex.Unlock()
	handlerInfoMap[info.ID] = info
}

func Handler(id string, middeName ...string) endpoint.Endpoint {
	mid := "bingo.default_middleware"
	if len(middeName) > 0 {
		mid = middeName[0]
	}
	if middle := GetMiddle(mid); middle != nil {
		return GetMiddle(mid)(handle(id))
	}
	return nil
}

func RunHandler(id string, req *Request) (interface{}, error) {
	h := Middleware()(handle(id))
	return h(context.Background(), req)
}

//---------------
func getHandler(id string) IHandler {
	if h, ok := handlerInfoMap[id]; ok {
		return h.New()
	}
	return nil
}
func handle(id string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		handler := getHandler(id)
		if handler == nil {
			return nil, fmt.Errorf("no handler for %s", id)
		}
		//生成请求参数,在编解码codec中变换的request
		req, _ := request.(*Request)
		//context
		newCtx := NewContext(ctx)
		//线程log,统一处理ip,request_id等
		newCtx.Log = makeLog(req)
		//设置request
		newCtx.Request = req
		//参数验证
		err = makeValid(newCtx, handler)
		if err != nil {
			newCtx.Log.Error(err.Error())
			return nil, err
		}
		//
		mock := req.GetString("mock")
		if mock == "yes" && viper.GetString("config.mode") == "dev" {
			if v, ok := handler.(IMock); ok {
				return v.Mock(newCtx), nil
			}
		}
		//逻辑处理
		ret, err := handler.Handle(newCtx)
		if err != nil {
			newCtx.Log.Error(err.Error())
		}
		return ret, err
	}
}

func makeValid(ctx Context, handler IHandler) error {
	// handler 的参数验证自动
	if v, ok := handler.(IValid); ok {
		if err := v.Valid(ctx); err != nil {
			return err
		}
	}
	return nil
}

func makeLog(req *Request) *logrus.Entry {
	//初始化日志字段,放到context中
	ip, has := (req.Data)["client_ip"]
	if !has || ip == nil {
		ip = "LAN"
	}
	entity := logger.GetInstance().WithFields(logrus.Fields{
		"request_id": req.Id,
		"client_ip":  ip,
	})
	return entity
}
