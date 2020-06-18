package test

import (
	"github.com/9299381/bingo/package/util"
	"testing"
)

type Src struct {
	Name  string
	Title string
}
type Dst struct {
	Name     string
	Password string
}

func TestStruct2Struct(t *testing.T) {
	src := &Src{
		Name:  "src_name",
		Title: "src_title",
	}
	dst := &Dst{}
	err := util.Struct2Struct(src, dst)
	if err != nil {
		t.Log(err)
	}
	t.Log(dst.Name)
}
