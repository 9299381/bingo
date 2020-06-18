package test

import (
	"fmt"
	"github.com/9299381/bingo/package/util"
	"testing"
)

func TestReflect(t *testing.T) {
	people := &People{m: map[string]interface{}{}}
	_, err := util.RefCall(people, "Set", "1", "中文")
	if err != nil {
		fmt.Println(err)
	}
	ret, err := util.RefCall(people, "Get", "1")
	fmt.Println(ret[0].Interface().(string))
	fmt.Println(people)
}

type People struct {
	m map[string]interface{}
}

func (p *People) Set(id string, name string) {
	fmt.Println("reflect set....")
	p.m[id] = name
}
func (p *People) Get(id string) interface{} {
	return p.m[id]
}
