package command

import (
	"context"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/go-kit/kit/endpoint"
	"github.com/spf13/viper"
	"sync"
)

func init() {
	bingo.RegisterModule(new(Command))
}

type Command struct {
	bingo.Context
	handlers map[string]bingo.ICommonHandler
	mutex    sync.RWMutex
}

func (c *Command) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "command",
		Type:    bingo.MODULE_NONE,
		Version: "1.0.0",
		New: func() bingo.IModule {
			cmd := new(Command)
			cmd.ConfigModule()
			return cmd
		},
	}
}

func (c *Command) ConfigModule() {
	c.Context = bingo.NewContext(context.Background())
	c.handlers = map[string]bingo.ICommonHandler{}
}

func (c *Command) Route(name string, endpoint endpoint.Endpoint) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	handler := NewCommandHandler(endpoint)
	c.handlers[name] = handler
}
func (c *Command) Start(id string) error {
	c.Log.Infof("%s start ...", id)
	route := viper.GetString("route")
	if handler, ok := c.handlers[route]; ok {
		args := viper.GetString("args")
		ret, err := handler.ServeHandle(context.Background(), args)
		if err != nil {
			return err
		}
		c.Log.Infof("%s command result : %v ", route, ret)
		return nil
	}
	return fmt.Errorf("no route for %s", route)
}
func (c *Command) Stop(id string) {
	c.Log.Infof("%s stop ...", id)
}

var (
	_ bingo.IModule       = (*Command)(nil)
	_ bingo.IModuleServer = (*Command)(nil)
)
