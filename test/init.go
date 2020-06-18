package test

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.Set("config.mode", "dev")
	viper.Set("config.name", "test")
	viper.Set("registy", "127.0.0.1:8500")
}
