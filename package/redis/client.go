package redis

import (
	"fmt"
	"github.com/9299381/bingo/package/config"
	redigo "github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var pool *redigo.Pool
var once sync.Once

func Pool() *redigo.Pool {
	once.Do(func() {
		pool = initRedis()
	})
	return pool
}

func initRedis() *redigo.Pool {
	timeout := time.Duration(config.EnvInt("redis.timeout", 10)) * time.Second
	pool := &redigo.Pool{
		Dial: func() (conn redigo.Conn, e error) {
			conn, err := redigo.Dial(
				"tcp",
				config.EnvString("redis.uri", "127.0.0.1:6937"),
				redigo.DialPassword(config.EnvString("redis.auth", "password")),
				redigo.DialDatabase(config.EnvInt("redis.db", 0)),
				redigo.DialConnectTimeout(timeout),
				redigo.DialReadTimeout(timeout),
				redigo.DialWriteTimeout(timeout),
			)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			return
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     config.EnvInt("redis.max_idle", 5),
		MaxActive:   config.EnvInt("redis.max_active", 50),
		IdleTimeout: timeout,
		Wait:        true,
	}
	return pool
}

func Int(reply interface{}, err error) (int, error) {
	return redigo.Int(reply, err)
}
func Uint64(reply interface{}, err error) (uint64, error) {
	return redigo.Uint64(reply, err)
}
func Float64(reply interface{}, err error) (float64, error) {
	return redigo.Float64(reply, err)
}
func Float64s(reply interface{}, err error) ([]float64, error) {
	return redigo.Float64s(reply, err)
}
func String(reply interface{}, err error) (string, error) {
	return redigo.String(reply, err)
}
func StringMap(reply interface{}, err error) (map[string]string, error) {
	return redigo.StringMap(reply, err)
}
func Strings(reply interface{}, err error) ([]string, error) {
	return redigo.Strings(reply, err)
}
func Bool(reply interface{}, err error) (bool, error) {
	return redigo.Bool(reply, err)
}
func Bytes(reply interface{}, err error) ([]byte, error) {
	return redigo.Bytes(reply, err)
}

func NewScript(keyCount int, src string) *redigo.Script {
	return redigo.NewScript(keyCount, src)
}

//------------
func Lock(key, value, delay string) (bool, error) {
	script := NewScript(3, luaLockScript())
	return Bool(script.Do(Pool().Get(), key, value, delay))
}
func UnLock(key string) (bool, error) {
	script := NewScript(1, luaUnLockScript())
	return Bool(script.Do(Pool().Get(), key))
}
