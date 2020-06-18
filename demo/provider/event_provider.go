package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/event"
)

type EventProvider struct {
}

func (e *EventProvider) Boot() {
}

func (e *EventProvider) Register() {
	bingo.Bind("event", func(module bingo.IModule) error {
		mod := module.(*event.Server)
		mod.Route("demo.event", bingo.Handler("demo.two"))
		return nil
	})
}
