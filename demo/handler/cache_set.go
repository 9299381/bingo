package handler

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/cache"
)

func init() {
	bingo.RegisterHandler(new(CacheSet))
}

type CacheSet struct {
}

func (*CacheSet) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.cache_set",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			return new(CacheSet)
		},
	}
}

func (*CacheSet) Handle(ctx bingo.Context) (interface{}, error) {
	v := make(map[string]interface{})
	v["aaa"] = "bbb"

	dd := make(map[string]interface{})
	dd["a"] = "111"
	dd["b"] = "222"

	v["ccc"] = dd
	_ = cache.Set("key", v, 60)

	return nil, nil
}

var (
	_ bingo.IHandler = (*CacheSet)(nil)
)
