package test

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/cronjob"
	"testing"
)

func TestModuleHelp(t *testing.T) {
	mod := bingo.Module("bingo.cron.server").(*cronjob.Server)
	ret := mod.Help()
	fmt.Println(ret)
}
