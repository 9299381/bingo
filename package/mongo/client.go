package mongo

import (
	"github.com/9299381/bingo/package/config"
	"github.com/9299381/bingo/package/logger"
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)

var session *mgo.Session
var once sync.Once

func Get() *mgo.Session {
	once.Do(func() {
		session = initMongo()
	})
	return session
}

func initMongo() *mgo.Session {
	idle := config.EnvInt("mongo.max_idle_time", 30)
	info := &mgo.DialInfo{
		Addrs:     config.EnvStringSlice("mongo.uri", []string{"127.0.0.1:27017"}),
		Timeout:   time.Duration(idle) * time.Second,
		Username:  config.EnvString("mongo.username", ""),
		Password:  config.EnvString("mongo.password", ""),
		Database:  config.EnvString("mongo.database", "base"),
		PoolLimit: config.EnvInt("mongo.min_pool_size", 10),
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		logger.GetInstance().Fatalf("MongoCreateSession: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
func Col(collection string, f func(*mgo.Collection)) {
	DB(config.EnvString("mongo.database", "bingo"), collection, f)
}

func DB(database string, collection string, f func(*mgo.Collection)) {
	session := Get().Clone()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			logger.GetInstance().Fatalf("MongoSessionError: %v", err)
		}
	}()
	col := session.DB(database).C(collection)
	f(col)
}
