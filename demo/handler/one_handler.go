package handler

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/http"
	"github.com/9299381/bingo/package/config"
)

func init() {
	bingo.RegisterHandler(new(OneHandler))
}

type OneHandler struct {
	Name string
}

func (*OneHandler) Info() *bingo.HandlerInfo {
	return &bingo.HandlerInfo{
		ID:      "demo.one",
		Version: "1.0.0",
		New: func() bingo.IHandler {
			handler := new(OneHandler)
			handler.ConfigHandler()
			return handler
		},
	}
}

func (*OneHandler) Handle(ctx bingo.Context) (interface{}, error) {

	//测试运行中增加,ok
	mod := bingo.Module("http").(*http.Server)
	mod.Get("/demo/add_two", bingo.Handler("demo.two", "comm"))
	//
	fmt.Println("one handler")
	m := make(map[string]interface{})
	m["handler"] = "one_handler"
	m["request"] = ctx.Request.Data
	return m, nil
}

func (h *OneHandler) ConfigHandler() {
	h.Name = config.EnvString("some.key", "default")
}

var (
	_ bingo.IHandler = (*OneHandler)(nil)
)
