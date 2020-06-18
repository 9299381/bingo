package test

import (
	"context"
	"fmt"
	"github.com/9299381/bingo"
	"testing"
)

func OneService(ctx bingo.Context) error {
	fmt.Println("one service")
	ctx.Set("one", "one")
	fmt.Println(ctx.GetString("bingo"))
	fmt.Println(ctx.GetString("two"))
	return nil
}
func TwoService(ctx bingo.Context) error {
	fmt.Println("two service")
	ctx.Set("two", "two")
	fmt.Println(ctx.GetString("bingo"))
	fmt.Println(ctx.GetString("one"))
	return nil
}

func TestLineService(t *testing.T) {
	ctx := bingo.NewContext(context.Background())
	ctx.Set("bingo", "bingo")
	err := bingo.Pipe(
		OneService,
		TwoService,
	).Line(ctx)
	fmt.Println(err)
}
func TestParallelService(t *testing.T) {
	ctx := bingo.NewContext(context.Background())
	ctx.Set("bingo", "bingo")
	_ = bingo.Pipe(
		OneService,
		TwoService,
	).Parallel(ctx)
	//fmt.Println(ctx.GetString("one"))
	//fmt.Println(ctx.GetString("two"))

}
