package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&server, "server", "s", "", "http,grpc,cron,event,mqtt,queue,registy,swagger,delayQ")
	startCmd.Flags().StringVarP(&registy, "registy", "r", "127.0.0.1:8500", "consul registy")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "后台运行run命令",
	//Long: "start command",
	//Args :cobra.MinimumNArgs(1), // 参数验证,可自定义func
	Run: func(command *cobra.Command, args []string) {
		//重复参数解析最后一个
		os.Args = append(os.Args, "--mode")
		os.Args = append(os.Args, "prod")
		cmd := exec.Command(os.Args[0], "run")
		for _, v := range os.Args[2:] {
			cmd.Args = append(cmd.Args, v)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("bingo start, [PID] %d running...\n", cmd.Process.Pid)
		_ = ioutil.WriteFile("bingo.lock", []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0666)
		os.Exit(0)
	},
}
