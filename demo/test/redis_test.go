package test

import (
	"fmt"
	"github.com/9299381/bingo/package/redis"
	"testing"
)

func TestRedisSet(t *testing.T) {
	conn := redis.Pool().Get()
	defer conn.Close()
	ret, _ := redis.Bool(conn.Do("HSET", "18605360126_sms", "sms", "1234"))
	fmt.Println(ret)
	_, _ = conn.Do("EXPIRE", "18605360126_sms", "20")
}
func TestRedisGetSet(t *testing.T) {
	conn := redis.Pool().Get()
	defer conn.Close()
	ret, _ := redis.String(conn.Do("HGET", "18605360126_sms", "sms"))
	fmt.Println(ret)
}

func TestLockScript(t *testing.T) {
	key := "123123123123"
	value := "trans_lock"
	delay := "60"
	ret, err := redis.Lock(key, value, delay)
	fmt.Println(ret, "------", err)
}
func TestUnLockScript(t *testing.T) {
	key := "123123123123"
	ret, err := redis.UnLock(key)
	fmt.Println(ret, "------", err)
}
