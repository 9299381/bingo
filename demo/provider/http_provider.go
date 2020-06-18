package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/http"
)

type HttpProvider struct {
}

func (h *HttpProvider) Boot() {
}

func (h *HttpProvider) Register() {
	// 推荐使用这种模式,方式1可能存在server未注册情况
	//方式2,hook方式,在server start时 hook
	bingo.Bind("http", func(module bingo.IModule) error {
		mod := module.(*http.Server)

		//default middleware
		mod.Get("/demo/one", bingo.Handler("demo.one"))
		mod.Get("/demo/two", bingo.Handler("demo.two"))
		mod.Get("/demo/event", bingo.Handler("demo.event"))
		mod.Get("/demo/redis", bingo.Handler("demo.redis"))
		mod.Get("/demo/queue", bingo.Handler("demo.queue"))
		mod.Get("/demo/cache_set", bingo.Handler("demo.cache_set"))
		mod.Get("/demo/cache_get", bingo.Handler("demo.cache_get"))
		mod.Get("/demo/consul", bingo.Handler("demo.consul"))
		mod.Get("/demo/valid", bingo.Handler("demo.valid"))
		mod.Get("/demo/async", bingo.Handler("demo.async_handler"))

		//auth middleware
		mod.Get("/demo/auth", bingo.Handler("demo.auth", "auth"))

		return nil
	})
}
