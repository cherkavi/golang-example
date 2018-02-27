package main

import "github.com/sparrc/go-ping"
import "fmt"

func main() {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		fmt.Errorf("ping error: ", err)
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println("ping result : ", stats.Addr)
}
