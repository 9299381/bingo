package db

import (
	"fmt"
	"github.com/9299381/bingo/package/config"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"xorm.io/xorm"
)

var db *xorm.Engine
var once sync.Once

func Engine() *xorm.Engine {
	once.Do(func() {
		db = initXorm()
	})
	return db
}

func initXorm() *xorm.Engine {
	driver := config.EnvString("db.connection", "mysql")
	dataSource := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s"+"?charset=utf8mb4&collation=utf8mb4_unicode_ci",
		config.EnvString("db.username", "root"),
		config.EnvString("db.password", "root"),
		config.EnvString("db.host", "127.0.0.1"),
		config.EnvString("db.port", "3306"),
		config.EnvString("db.database", "default"),
	)
	engine, err := xorm.NewEngine(driver, dataSource)
	if err != nil {
		panic(err)
	}
	//连接池
	engine.SetMaxIdleConns(config.EnvInt("db.max_idle", 5))
	engine.SetMaxOpenConns(config.EnvInt("db.max_open", 50))
	engine.ShowSQL(config.EnvBool("db.show_sql", false))
	return engine
}
