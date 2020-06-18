package test

import (
	"fmt"
	"github.com/9299381/bingo/package/config"
	"github.com/spf13/viper"
	"testing"
)

func TestNewCfg(t *testing.T) {
	v := viper.New()
	v.SetDefault("aaaaa", "bbb")
	v.SetConfigName("app")
	v.AddConfigPath("./")
	v.SetConfigType("toml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))

	}
	fmt.Println(v.Get("aaaaa"))
	fmt.Println(v.Get("common.cache.size"))
}
func TestConfig(t *testing.T) {

	queue := config.Env("queue.listen", "")
	port := config.Env("server.http_port", "")
	fmt.Println(queue)
	fmt.Println(port)

}
