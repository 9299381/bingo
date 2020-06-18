package provider

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/modules/mqtt"
)

type MQTTProvider struct {
}

func (M *MQTTProvider) Boot() {
}

func (M *MQTTProvider) Register() {
	bingo.Bind("mqtt", func(module bingo.IModule) error {
		mod := module.(*mqtt.Subscribe)
		mod.Route("sub_test", bingo.Handler("demo.one"))
		mod.Route("$SYS/brokers/+/clients/+/+",
			bingo.Handler("demo.mqtt_event"))

		return nil
	})
}
