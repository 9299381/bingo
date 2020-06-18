package cmd

import (
	"github.com/9299381/bingo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	server  string
	registy string
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&server, "server", "s", "", "http,grpc,cron,event,mqtt,queue,registy,swagger,delayQ")
	runCmd.Flags().StringVarP(&registy, "registy", "r", "127.0.0.1:8500", "consul registy")
	_ = viper.BindPFlags(runCmd.Flags())

}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "前台运行",
	//Long: "run command",
	Run: func(cmd *cobra.Command, args []string) {
		// 先定义退出 chan
		quit := make(chan error)
		bingo.Signals(bingo.ModuleStop)
		// 各模块初始化加载
		bingo.ModuleLoad(server)
		// 注入的 provider 运行
		bingo.Provide()
		// 模块启动
		bingo.ModuleStart(quit)
		// 模块停止
		bingo.ModuleStop(<-quit)
	},
}
