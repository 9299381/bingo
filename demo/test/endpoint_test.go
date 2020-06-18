package test

import (
	"context"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
	"testing"
)

func TestEndpoint(t *testing.T) {

	middle := bingo.Middleware(
		testOneMiddle(),
		testTwoMiddle(),
	)
	bingo.RegisterMiddleware("middle", middle)

	e1 := bingo.Handler("demo.two", "middle")
	ctx := bingo.NewContext(context.Background())

	req := &bingo.Request{
		Id: "123",
		Data: map[string]interface{}{
			"aa": "bb",
		},
	}
	ret, err := e1(ctx, req)
	if err == nil {
		fmt.Println(ret)
	}

}

func testOneMiddle() endpoint.Middleware {
	fmt.Println("go test 1")
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		fmt.Println("do test 1")
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			fmt.Println("do handle 1")
			return next(ctx, request)
		}
	}
}
func testTwoMiddle() endpoint.Middleware {
	fmt.Println("go test 2")
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		fmt.Println("do test 2")
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			fmt.Println("do handle 2")
			return next(ctx, request)
		}
	}
}
