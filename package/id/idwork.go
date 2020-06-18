package id

import (
	"github.com/9299381/bingo/package/config"
	"strconv"
	"sync"
	"time"
)

var worker *IdWork
var once sync.Once

func instance() *IdWork {
	once.Do(func() {
		worker = initID()
	})
	return worker
}

func New() string {
	return instance().GetID()
}

type IdWork struct {
	mu           sync.Mutex
	lastTime     int64
	server       int64
	serverMax    int64
	sequence     int64
	sequenceMask int64

	timeShift   uint8
	serverShift uint8
}

func initID() *IdWork {
	serverId := config.EnvInt("server_id", 512)
	it := &IdWork{}
	it.server = int64(serverId)
	var serverBits uint8 = 10
	var sequenceBits uint8 = 12
	it.serverMax = -1 ^ (-1 << serverBits)
	it.sequenceMask = -1 ^ (-1 << sequenceBits)
	it.serverShift = sequenceBits
	it.timeShift = serverBits + sequenceBits
	if it.server < 0 || it.server > it.serverMax {
		panic("server number must be between 0 and " + strconv.FormatInt(it.serverMax, 10))
	}
	return it
}

func (it *IdWork) GetID() string {
	it.mu.Lock()
	epoch := it.epochGen()
	timestamp := time.Now().UnixNano() / 1000000
	lastTime := it.lastTime
	//生成唯一序列
	if timestamp == lastTime {
		it.sequence = (it.sequence + 1) & it.sequenceMask
		if it.sequence == 0 {
			for timestamp <= lastTime {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		it.sequence = 0
	}
	it.lastTime = timestamp
	r := (timestamp-epoch)<<it.timeShift |
		(it.server << it.serverShift) |
		(it.sequence)
	it.mu.Unlock()
	return strconv.FormatInt(int64(r), 10)
}

func (it *IdWork) epochGen() int64 {
	start := "2010-01-01 00:00:00"
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, start, loc)
	return theTime.UnixNano() / 1000000
}
