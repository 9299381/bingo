package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
	"testing"
)

func TestHandler(t *testing.T) {
	req := &bingo.Request{
		Id: id.New(),
		Data: map[string]interface{}{
			"hello": "world",
		},
	}
	ret, err := bingo.RunHandler("demo.two", req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
