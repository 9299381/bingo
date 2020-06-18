package service

import "github.com/9299381/bingo"

func OneService(ctx bingo.Context) error {
	ctx.Set("aaa", "aaa")
	return nil
}
func TwoService(ctx bingo.Context) error {
	ctx.Log.Info(ctx.Get("aaa"))
	ctx.Set("bbb", "bbb")
	return nil
}
