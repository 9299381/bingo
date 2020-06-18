package test

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	ret := getRate()
	fmt.Println(ret)
}

func getRate() int {
	num1 := 65
	num2 := 1000
	ret := num1 * num2 / 10000
	return ret

}
