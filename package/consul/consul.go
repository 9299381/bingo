package consul

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/9299381/bingo/package/cache"
	"github.com/9299381/bingo/package/id"
	"github.com/9299381/bingo/package/logger"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
)

func RegistyHttp(service, host, port string) *consul.Registrar {
	//注销掉重复的
	deregister(service, host, port)
	check := api.AgentServiceCheck{
		HTTP:     "http://" + host + ":" + port + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	p, _ := strconv.Atoi(port)
	reg := api.AgentServiceRegistration{
		ID:      service + "_" + id.New(),
		Name:    service,
		Address: host,
		Port:    p,
		Tags:    []string{"http"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(getClient(), &reg, newKitLog())
	registy.Register()
	return registy
}

func RegistyGrpc(service, host, port string) *consul.Registrar {
	//注销掉重复的
	deregister(service, host, port)
	p, _ := strconv.Atoi(port)
	check := api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d/%s", host, p, "health"),
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	reg := api.AgentServiceRegistration{
		ID:      service + "_" + id.New(),
		Name:    service,
		Address: host,
		Port:    p,
		Tags:    []string{"grpc"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(getClient(), &reg, newKitLog())
	registy.Register()
	return registy

}

func GetService(service string) (entity *api.ServiceEntry, err error) {
	//这里考虑可以从缓存中读取,10分钟过期,比如
	var entitys []*api.ServiceEntry
	c, _ := cache.GetByte("consul_entitys")
	if c != nil {
		entitys = []*api.ServiceEntry{}
		err = json.Unmarshal(c, &entitys)
		if err != nil {
			panic(err)
			return
		}
	} else {
		client := getClient()
		entitys, _, err = client.Service(service, "", false, &api.QueryOptions{})
		if err != nil || len(entitys) == 0 {
			err = errors.New("9999::没有找到响应的服务")
			return
		}
		_ = cache.Set("consul_entitys", entitys, 60)
	}
	//随机取一个
	entity = entitys[rand.Int()%len(entitys)]
	return

}

func getClient() consul.Client {
	var client consul.Client
	{
		config := api.DefaultConfig()
		config.Address = viper.Get("registy").(string)
		consulClient, _ := api.NewClient(config)
		client = consul.NewClient(consulClient)
	}
	return client
}

func deregister(service, host, port string) {
	client := getClient()
	entitys, _, err := client.Service(service, "", false, &api.QueryOptions{})
	if err == nil {
		for _, entity := range entitys {
			str1 := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
			str2 := fmt.Sprintf("%s:%s", host, port)
			if str1 == str2 {
				r := &api.AgentServiceRegistration{
					ID:   entity.Service.ID,
					Name: service,
				}
				_ = client.Deregister(r)
			}
		}
	}
}

type goLogger struct {
	*logrus.Logger
}

func (s goLogger) Log(args ...interface{}) error {
	s.Info(args)
	return nil
}

// log.logger 接口的一种实现,用以注入 go-kit 的服务注册
func newKitLog() log.Logger {
	return goLogger{
		Logger: logger.GetInstance(),
	}
}
