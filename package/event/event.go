package event

import (
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/event"
)

func Fire(payload *bingo.Payload) error {
	if mod, ok := bingo.Module("event").(*event.Server); ok {
		route := mod.GetRoute(payload.Route)
		if route != nil {
			return mod.AddEvent(payload)
		} else {
			return fmt.Errorf("event no handler:%s", payload.Route)
		}
	}
	return fmt.Errorf("%s register error", "event")
}
