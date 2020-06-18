package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/micro"
	"testing"
)

func TestMicro(t *testing.T) {
	ret := micro.Service("demo").
		Api("demo.two").
		Request(map[string]interface{}{
			"aaa":    "123123",
			"bbbbbb": "ok",
		}).Run().Data.(*bingo.Response)
	fmt.Println(ret)
	fmt.Println(ret.Data)
}
