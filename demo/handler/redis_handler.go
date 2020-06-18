package handler

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/redis"
)

func init() {
	bingo.RegisterHandler(new(RedisHandler))
}

type RedisHandler struct {
}

func (*RedisHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.redis",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(RedisHandler)
		},
	}
}

func (*RedisHandler) Handle(ctx bingo.Context) (interface{}, error) {
	conn := redis.Pool().Get()
	defer conn.Close()
	_, _ = conn.Do("SET", "go_key", "value")
	res, _ := redis.String(conn.Do("GET", "go_key"))
	exists, _ := redis.Bool(conn.Do("EXISTS", "foo"))
	if exists {
		ctx.Log.Info("foo 存在")
	} else {
		_, _ = conn.Do("SET", "foo", "value")
		ctx.Log.Info("foo 不存在")

	}
	ctx.Log.Info("redis-go_key 的值:", res)
	return res, nil
}

var (
	_ bingo.IHandler = (*RedisHandler)(nil)
)
