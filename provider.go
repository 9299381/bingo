package bingo

import (
	"sync"
)

var (
	providers []IProvider
	mutex     sync.RWMutex
)

func Provider(provider IProvider) {
	mutex.Lock()
	defer mutex.Unlock()
	providers = append(providers, provider)
}

func getProviders() []IProvider {
	mutex.RLock()
	defer mutex.RUnlock()
	return providers
}

func Provide() {
	for _, provider := range getProviders() {
		provider.Boot()
		provider.Register()
	}
}
