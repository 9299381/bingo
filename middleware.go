package bingo

import (
	"context"
	"errors"
	"fmt"
	"github.com/9299381/bingo/package/logger"
	"github.com/9299381/bingo/package/redis"
	"github.com/9299381/bingo/package/token"
	"github.com/9299381/bingo/package/util"
	"github.com/go-kit/kit/endpoint"
	"strings"
	"sync"
)

var (
	middleMap      = make(map[string]endpoint.Middleware)
	middleMapMutex sync.Mutex
)

func init() {
	RegisterMiddleware("bingo.default_middleware", Middleware())
}
func RegisterMiddleware(id string, m endpoint.Middleware) {
	if _, ok := middleMap[id]; ok {
		panic(fmt.Sprintf("middleware already registered: %s", id))
	}
	middleMapMutex.Lock()
	defer middleMapMutex.Unlock()
	middleMap[id] = m
}
func GetMiddle(id string) endpoint.Middleware {
	if m, ok := middleMap[id]; ok {
		return m
	}
	return Middleware()
}

//
func Middleware(middles ...endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		for i := len(middles) - 1; i >= 0; i-- { // reverse
			next = middles[i](next)
		}
		return ResponseMiddleware()(next)
	}
}

// 响应
func ResponseMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			//全局扑捉错误
			defer func() {
				if err := recover(); err != nil {
					logger.GetInstance().Info(err)
					response = MakeResponse(nil, err.(error))
				}
			}()
			response, err = next(ctx, request)
			return MakeResponse(response, err), nil
		}
	}
}

// token auth
func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(*Request)
			auth := req.Data["authToken"]
			if auth == nil || auth == "" {
				return nil, errors.New(token.ErrNoToken)
			}
			claim, err := token.CheckToken(auth.(string))
			if err != nil {
				return nil, err
			}
			req.Data["claim"] = util.Struct2Map(claim)
			//这里进行token的jwt认证
			return next(ctx, req)
		}
	}
}

// 交易锁定中间件
func LockMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(*Request)
			if userId := req.Claim("id"); userId != "" {
				key := strings.Join([]string{"trans", userId}, "_")
				//lock user
				lock, err := redis.Lock(key, "lock", "60")
				if err != nil {
					return nil, err
				}
				if lock {
					defer func() {
						_, _ = redis.UnLock(key)
					}()
					return next(ctx, request)
				}
				return nil, fmt.Errorf("waiting transaction for %s", userId)
			}
			return nil, fmt.Errorf("must auth middleware before")
		}
	}
}
