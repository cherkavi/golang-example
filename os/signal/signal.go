package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ExampleOfSignalHook() {
	fmt.Println(" > start signal listener - waiting for Ctrl-C ")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	signal := <-ch
	fmt.Println("signal from OS: ", signal)
}
