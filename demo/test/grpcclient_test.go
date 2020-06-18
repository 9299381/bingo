package test

import (
	"fmt"
	"github.com/9299381/bingo/package/grpc"
	"testing"
)

func TestGrpcClient(t *testing.T) {
	host := "127.0.0.1:9341"
	service := "demo.one"
	m := make(map[string]interface{})
	m["one"] = "one"
	m["two"] = "two"
	ret := grpc.Call(host, service, m)
	fmt.Println(ret)

}
