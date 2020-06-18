package mqtt

import (
	"encoding/json"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/id"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(topic string, payload map[string]interface{}) error {
	param, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	qos := uint8(config.EnvInt("mqtt.publish_qos", 2))
	client := initClient()
	token := client.Publish(topic, qos, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func initClient() mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(config.EnvString("mqtt.host", "tcp://127.0.0.1:1883"))
	opts.SetUsername(config.EnvString("mqtt.username", ""))
	opts.SetPassword(config.EnvString("mqtt.password", ""))
	opts.SetClientID(config.EnvString("mqtt.clientid", id.New()))
	mc := mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return mc
}
