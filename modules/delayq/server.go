package delayq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/redis"
	"strings"
	"time"
)

func init() {
	bingo.RegisterModule(new(DelayServer))
}

type DelayServer struct {
	bingo.Context
	freq      int
	delayName string
	prefix    string
}

func (*DelayServer) ModuleInfo() *bingo.ModuleInfo {
	return &bingo.ModuleInfo{
		ID:      "delayQ",
		Type:    bingo.MODULE_SERVER,
		Version: "1.0.0",
		New: func() bingo.IModule {
			ss := new(DelayServer)
			ss.ConfigModule()
			return ss
		},
	}
}

func (s *DelayServer) ConfigModule() {
	s.Context = bingo.NewContext(context.Background())
	s.freq = config.EnvInt("queue.interval", 2)
	s.prefix = config.EnvString("queue.prefix", "bingo")
	s.delayName = strings.Join([]string{s.prefix, "queue_delay"}, "_")
}
func (s *DelayServer) Start(id string) error {
	conn := redis.Pool().Get()
	defer conn.Close()

	errChan := make(chan error)
	ticker := time.NewTicker(time.Duration(s.freq) * time.Second)
	go func(ch chan error, t *time.Ticker) {
		for {
			select {
			case <-t.C:
				now := time.Now().Unix()
				reply, err := redis.Strings(conn.Do("ZRANGEBYSCORE", s.delayName, 0, now))
				if err != nil {
					ch <- err
				}
				if len(reply) > 0 {
					for _, v := range reply {
						job := &bingo.Job{}
						s.Log.Infof("found delay job on %s", job.Queue)
						err := json.Unmarshal([]byte(v), job)
						if err != nil {
							ch <- err
						}
						key := fmt.Sprintf("%s_queue:%s", s.prefix, job.Queue)
						_ = conn.Send("RPUSH", key, []byte(v))
						_ = conn.Send("ZREM", s.delayName, v)
					}
					_ = conn.Flush()
				}

			}
		}
	}(errChan, ticker)

	return <-errChan
}

func (s *DelayServer) Stop(id string) {
	s.Log.Infof("%s stop ...", id)
}

var (
	_ bingo.IModule       = (*DelayServer)(nil)
	_ bingo.IModuleServer = (*DelayServer)(nil)
)
