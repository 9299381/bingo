package cache

import (
	"encoding/json"
	"github.com/9299381/bingo/package/config"
	"github.com/coocood/freecache"
	"runtime/debug"
	"sync"
)

var once sync.Once
var ins *freecache.Cache

func instance() *freecache.Cache {
	once.Do(func() {
		ins = initCache()
	})
	return ins
}

func initCache() *freecache.Cache {
	size := config.EnvInt("cache.size", 1048576)
	if size != 0 {
		c := freecache.NewCache(size)
		//根据cache的大小进行设置
		debug.SetGCPercent(20)
		return c
	}
	return nil
}

func Set(key string, value interface{}, exp int) error {
	k := []byte(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = instance().Set(k, v, exp)
	if err != nil {
		return err
	}
	return nil
}
func Get(key string, obj interface{}) error {
	b, err := GetByte(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}
func GetByte(key string) ([]byte, error) {
	k := []byte(key)
	return instance().Get(k)
}
