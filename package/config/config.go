package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitConfig(cfgFile string) {
	if strings.Contains(cfgFile, ":") {
		loadFromConsul(cfgFile)
	} else {
		loadFromToml(cfgFile)
	}
}

func loadFromToml(cfgFile string) {
	if exist, _ := pathExists(cfgFile); exist {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("../")
		viper.AddConfigPath("../bin/")
		viper.AddConfigPath("./bin/")
	}
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file not found")
		os.Exit(0)
	}

}
func loadFromConsul(url string) {

}

func Env(key string, value interface{}) interface{} {
	modeKey := strings.Join([]string{viper.GetString("config.mode"), key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")

	if viper.IsSet(modeKey) {
		return viper.Get(modeKey)
	} else if viper.IsSet(commKey) {
		return viper.Get(commKey)
	} else {
		return value
	}
}
func EnvString(key string, value string) string {
	modeKey := strings.Join([]string{viper.GetString("config.mode"), key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret string
	if viper.IsSet(modeKey) {
		ret = viper.GetString(modeKey)
	} else if viper.IsSet(commKey) {
		ret = viper.GetString(commKey)
	} else {
		ret = value
	}
	return ret
}
func EnvInt(key string, value int) int {
	modeKey := strings.Join([]string{viper.GetString("config.mode"), key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret int
	if viper.IsSet(modeKey) {
		ret = viper.GetInt(modeKey)
	} else if viper.IsSet(commKey) {
		ret = viper.GetInt(commKey)
	} else {
		ret = value
	}
	return ret
}
func EnvBool(key string, value bool) bool {
	modeKey := strings.Join([]string{viper.GetString("config.mode"), key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret bool
	if viper.IsSet(modeKey) {
		ret = viper.GetBool(modeKey)
	} else if viper.IsSet(commKey) {
		ret = viper.GetBool(commKey)
	} else {
		ret = value
	}
	return ret
}
func EnvStringSlice(key string, value []string) []string {
	modeKey := strings.Join([]string{viper.GetString("config.mode"), key}, ".")
	commKey := strings.Join([]string{"common", key}, ".")
	var ret []string
	if viper.IsSet(modeKey) {
		ret = viper.GetStringSlice(modeKey)
	} else if viper.IsSet(commKey) {
		ret = viper.GetStringSlice(commKey)
	} else {
		ret = value
	}
	return ret
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
