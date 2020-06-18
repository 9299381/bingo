package test

import (
	"context"
	"errors"
	"fmt"
	"github.com/9299381/bingo"
	"testing"
)

func TestDefer(t *testing.T) {
	ctx := bingo.NewContext(context.Background())
	ctx.Set("start", "start")

	defer func(dtx bingo.Context) {
		fmt.Println(dtx.GetString("after"))
	}(ctx)

	if err := demo(ctx); err != nil {
		fmt.Println(err)
	}
}

func demo(ctx bingo.Context) error {
	ctx.Set("after", "after")
	return errors.New("err")
}
