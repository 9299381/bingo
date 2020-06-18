package mqtt

import (
	"context"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/id"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/endpoint"
	"sync"
)

func init() {
	bingo.RegisterModule(new(Subscribe))
}

type Subscribe struct {
	mqtt.Client
	bingo.Context

	handlers map[string]bingo.ICommonHandler
	mutex    sync.RWMutex

	Parallel     bool //并行处理
	SubscribeQos byte
}

func (s *Subscribe) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "mqtt",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0.0",
		New: func() bingo.IModule {
			server := new(Subscribe)
			server.ConfigModule()
			return server
		},
	}
}

func (s *Subscribe) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	opts := mqtt.NewClientOptions().AddBroker(config.EnvString("mqtt.host", "tcp://127.0.0.1:1883"))
	opts.SetUsername(config.EnvString("mqtt.username", ""))
	opts.SetPassword(config.EnvString("mqtt.password", ""))
	opts.SetClientID(config.EnvString("mqtt.clientid", id.New()))
	mc := mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	s.Client = mc
	s.handlers = map[string]bingo.ICommonHandler{}
	s.Parallel = config.EnvBool("mqtt.parallel", false)
	s.SubscribeQos = uint8(config.EnvInt("mqtt.subscribe_qos", 2))
}

func (s *Subscribe) Route(name string, endpoint endpoint.Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	handler := NewMqttSubscribe(endpoint)
	s.handlers[name] = handler
}

func (s *Subscribe) getHandler(name string) bingo.ICommonHandler {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.handlers[name]
}

func (s *Subscribe) Start(id string) error {
	s.Log.Infof("%s start ...", id)
	errChans := make(map[string]chan error)
	s.work(errChans)
	for _, errChan := range errChans {
		if errChan != nil {
			s.Log.Info(<-errChan)
		}
	}
	return nil
}
func (s *Subscribe) Stop(id string) {
	s.Log.Infof("%s stop ...", id)
}
func (s *Subscribe) work(errChans map[string]chan error) {
	s.Log.Info("MQTT Subscribe Server Start")
	for topic, handler := range s.handlers {
		errChans[topic] = make(chan error)
		go s.worker(topic, handler, errChans[topic])
	}

}
func (s *Subscribe) worker(t string, h bingo.ICommonHandler, e chan error) {
	s.Log.Infof("Subscribe topic:%s", t)
	token := s.Subscribe(t, s.SubscribeQos, func(
		client mqtt.Client, message mqtt.Message) {
		if s.Parallel {
			go s.process(h, message)
		} else {
			s.process(h, message)
		}
	})
	if token.Wait() && token.Error() != nil {
		e <- token.Error()
	}
}
func (s *Subscribe) process(h bingo.ICommonHandler, Message mqtt.Message) {
	s.Log.Info("subscribe topic:", Message.Topic())
	resp, err := h.ServeHandle(context.Background(), Message)
	if err != nil {
		s.Log.Error(err)

	} else {
		s.Log.Info(resp)
	}
}

var (
	_ bingo.IModule       = (*Subscribe)(nil)
	_ bingo.IModuleServer = (*Subscribe)(nil)
)
