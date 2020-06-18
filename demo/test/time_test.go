package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/util"
	"testing"
)

func TestTime(t *testing.T) {
	dateTime := util.GetLastMonth()
	fmt.Println(dateTime.Start.Format(bingo.YmdHis))
	fmt.Println(dateTime.End.Format(bingo.YmdHis))

}
