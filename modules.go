package bingo

import (
	"fmt"
	"github.com/9299381/bingo/package/logger"
	"os"
	"strings"
	"sync"
)

var (
	moduleInfoMap = make(map[string]*ModuleInfo)
	moduleMap     = make(map[string]IModule)
	moduleHook    = make(map[string][]func(IModule) error)
	moduleLock    sync.Mutex
)

const (
	MODULE_BOOT   = 0      //引导式插件
	MODULE_NONE   = 1      //独立插件
	MODULE_HOOK   = 1 << 1 //钩子插件
	MODULE_SERVER = 1 << 2 //应用插件
)

//模块必须实现的接口
type IModule interface {
	ModuleInfo() *ModuleInfo
}

type IModuleServer interface {
	Start(id string) error
	Stop(id string)
}

// 模块info
type ModuleInfo struct {
	ID      string
	Type    byte   //类型
	Version string //版本,根据版本高低,替换掉注册,todo
	New     func() IModule
}

// 注册模块
func RegisterModule(mod IModule) {
	info := mod.ModuleInfo()
	if info.ID == "" {
		panic("module ID missing")
	}
	if info.New == nil {
		panic("missing ModuleInfo.New")
	}
	if _, ok := moduleInfoMap[info.ID]; ok {
		panic(fmt.Sprintf("module already registered: %s", info.ID))
	}
	moduleLock.Lock()
	defer moduleLock.Unlock()
	moduleInfoMap[info.ID] = info
}

func GetModuleInfo(id string) (*ModuleInfo, error) {
	if info, ok := moduleInfoMap[id]; ok {
		return info, nil
	}
	return nil, fmt.Errorf("module not registered: %s", id)
}

func GetModules() map[string]IModule {
	return moduleMap
}
func Module(id string) IModule {
	if mod, ok := moduleMap[id]; ok {
		return mod
	}
	return nil
}

func Bind(id string, f func(IModule) error) {
	if _, err := GetModuleInfo(id); err == nil {
		moduleHook[id] = append(moduleHook[id], f)
	}
}

//module server start,直接在传入的server
func ModuleLoad(server string) {
	array := strings.Split(server, ",")
	moduleLock.Lock()
	defer moduleLock.Unlock()
	for _, v := range array {
		if info, err := GetModuleInfo(v); err == nil {
			moduleMap[info.ID] = info.New()
		}
	}
}

func GetModuleHooker(key string) []func(IModule) error {
	return moduleHook[key]
}

func ModuleStart(quit chan error) {
	for id, mod := range GetModules() {
		if mod.ModuleInfo().Type >= MODULE_HOOK {
			go func(key string, module IModule) {
				//hook
				for _, hooker := range GetModuleHooker(key) {
					if err := hooker(module); err != nil {
						quit <- err
					}
				}
				//server
				if server, ok := module.(IModuleServer); ok {
					if err := server.Start(key); err != nil {
						quit <- err
					}
				}
			}(id, mod)
		}
	}
}
func ModuleStop(err error) {
	for id, mod := range GetModules() {
		if mod.ModuleInfo().Type == MODULE_SERVER {
			if server, ok := mod.(IModuleServer); ok {
				server.Stop(id)
			}
		}
	}
	logger.GetInstance().Infof("bingo stop for:%s", err.Error())
	os.Exit(0)
}
