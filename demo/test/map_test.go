package test

import (
	"context"
	"fmt"
	"github.com/9299381/bingo"
	"testing"
)

func TestMap(t *testing.T) {

	ctx := bingo.NewContext(context.Background())
	m := make(map[string]interface{})
	m["aa"] = "1"
	m["bb"] = 2
	ctx.Set("a.b", "aaa")
	ctx.Set("a.c", "bbb")
	ctx.Set("a.d.e", m)

	ret := ctx.GetInt("a.d.e.bb")
	fmt.Println(ret)

}
