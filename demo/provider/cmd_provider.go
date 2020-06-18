package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/command"
)

type CmdProvider struct{}

func (CmdProvider) Boot() {

}

func (CmdProvider) Register() {
	bingo.Bind("command", func(module bingo.IModule) error {
		mod := module.(*command.Command)
		mod.Route("demo.cmd", bingo.Handler("demo.cmd"))
		return nil
	})
}
