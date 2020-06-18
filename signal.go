package bingo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signals(f func(err error)) {
	//创建监听退出chan
	ch := make(chan os.Signal)
	//监听指定信号 ctrl+c kill.....
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for sig := range ch {
			f(fmt.Errorf("quit for %s", sig.String()))
		}
	}()
}
