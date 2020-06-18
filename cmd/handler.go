package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/9299381/bingo"
	"github.com/9299381/bingo/package/id"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	rootCmd.AddCommand(handlerCmd)
	handlerCmd.Flags().StringVarP(&route, "route", "r", "", "require route")
	handlerCmd.Flags().StringVarP(&args, "args", "a", "", `example: '{"key":"value"}'`)
	_ = handlerCmd.MarkFlagRequired("route")
	_ = viper.BindPFlags(handlerCmd.Flags())
}

var handlerCmd = &cobra.Command{
	Use:   "handle",
	Short: "命令行运行handler",
	Run: func(cmd *cobra.Command, args []string) {
		data := make(map[string]interface{})
		jsonStr := viper.GetString("args")
		if jsonStr != "" {
			err := json.Unmarshal([]byte(jsonStr), &data)
			if err != nil {
				fmt.Println("args参数json解析错误")
				os.Exit(0)
			}
		}
		params := &bingo.Request{Id: id.New(), Data: data}
		ret, err := bingo.RunHandler(route, params)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
		os.Exit(0)
	},
}
