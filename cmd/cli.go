package cmd

import (
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	route string
	args  string
)

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&route, "route", "r", "", "require route")
	cliCmd.Flags().StringVarP(&args, "args", "a", "", `example: '{"key":"value"}'`)
	_ = cliCmd.MarkFlagRequired("route")
	_ = viper.BindPFlags(cliCmd.Flags())
}

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "命令行运行cmd路由",
	//Long: "cli command",
	Run: func(cmd *cobra.Command, args []string) {
		id := "command"
		if info, err := bingo.GetModuleInfo(id); err == nil {
			mod := info.New()
			bingo.Provide()
			for _, hooker := range bingo.GetModuleHooker(id) {
				if err := hooker(mod); err != nil {
					logger.GetInstance().Error(err)
				}
			}
			if server, ok := mod.(bingo.IModuleServer); ok {
				if err := server.Start(id); err != nil {
					logger.GetInstance().Error(err)
				}
				server.Stop(id)
			}
		}
	},
	//运行前和运行后钩子
	//PersistentPreRun
	//PreRun
	//Run
	//PostRun
	//PersistentPostRun
}
