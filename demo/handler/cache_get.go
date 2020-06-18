package handler

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/cache"
	"github.com/tidwall/gjson"
)

func init() {
	bingo.RegisterHandler(new(CacheGet))
}

type CacheGet struct{}

func (*CacheGet) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.cache_get",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(CacheGet)
		},
	}
}

func (CacheGet) Handle(ctx bingo.Context) (interface{}, error) {
	//GetByte 方式
	jsonBytes, _ := cache.GetByte("key")
	// Get方式
	obj := make(map[string]interface{})
	_ = cache.Get("key", &obj)

	//使用gjson 更方便 直接从json字符串中取值
	return map[string]interface{}{
		"ccc":  gjson.Get(string(jsonBytes), "ccc.a").Str,
		"aaa":  obj["aaa"],
		"gaaa": gjson.Get(string(jsonBytes), "aaa").String(),
	}, nil
}

var (
	_ bingo.IHandler = (*CacheGet)(nil)
)
