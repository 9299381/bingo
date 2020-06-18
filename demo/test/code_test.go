package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	code := "K2100"
	ret := strings.Join([]string{
		code[0:3], "01",
	}, "")
	fmt.Println(ret)
}
