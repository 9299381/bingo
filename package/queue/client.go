package queue

import (
	"encoding/json"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/redis"
	"strings"
	"time"
)

func Fire(job *bingo.Job) error {
	conn := redis.Pool().Get()
	defer conn.Close()
	prefix := config.EnvString("queue.prefix", "bingo")
	buffer, err := json.Marshal(job)
	if err != nil {
		return err
	}
	err = conn.Send("RPUSH", fmt.Sprintf("%s_queue:%s", prefix, job.Queue), buffer)
	if err != nil {
		return err
	}
	return conn.Flush()
}
func Delay(delay int, job *bingo.Job) error {
	conn := redis.Pool().Get()
	defer conn.Close()
	execTime := time.Now().Add(time.Second * time.Duration(delay)).Unix()
	buffer, _ := json.Marshal(job)
	prefix := config.EnvString("queue.prefix", "bingo")
	key := strings.Join([]string{prefix, "queue_delay"}, "_")
	err := conn.Send("ZADD", key, execTime, buffer)
	if err != nil {
		return err
	}
	return conn.Flush()
}
func UnDelay(job *bingo.Job) error {
	conn := redis.Pool().Get()
	defer conn.Close()
	buffer, _ := json.Marshal(job)
	prefix := config.EnvString("queue.prefix", "bingo")
	queueKey := fmt.Sprintf("%s_queue:%s", prefix, job.Queue)
	delayKey := strings.Join([]string{prefix, "queue_delay"}, "_")
	_ = conn.Send("RPUSH", queueKey, buffer)
	_ = conn.Send("ZREM", delayKey, buffer)
	return conn.Flush()
}
