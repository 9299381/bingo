package main

import (
	"github.com/9299381/bingo/cmd"
	_ "github.com/9299381/bingo/demo/provider"
	_ "github.com/9299381/bingo/modules"
	"os"
)

func main() {
	//
	os.Args = append(os.Args, "run")
	os.Args = append(os.Args, "--config")
	os.Args = append(os.Args, "../../../config.toml")

	os.Args = append(os.Args, "--server")
	os.Args = append(os.Args, "http")

	//os.Args = append(os.Args,"http,queue,delayQ")
	//

	cmd.Execute()
}
