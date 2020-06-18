package cmd

import (
	"fmt"
	"github.com/9299381/bingo/package/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	mode    string
	name    string
)
var rootCmd = &cobra.Command{
	Use: "bingo",
	//Short: "short description",
	//Long: "long description",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./config.toml", "配置文件")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "bingo", "应用名称")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "运行模式")
	viper.Set("config.mode", &mode)
	viper.Set("config.name", &name)
}

//--------
func initConfig() {
	config.InitConfig(cfgFile)
	// todo 设置一下docker下ip
}
