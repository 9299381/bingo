package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止后台run命令",
	//Long: `stop command`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "bingo.lock"
		strb, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			_ = os.Remove(filename)
			command := exec.Command("kill", string(strb))
			_ = command.Start()
			fmt.Println("bingo stop")
		}
	},
}
