package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("  Sleep for period of time  ")
	time.Sleep(10 * time.Millisecond)

	fmt.Println("waiting via channel")
	select {
	case time := <-time.After(100 * time.Millisecond):
		{
			fmt.Printf("timeout: %#v \n", time)
			goto Next
		}
	}
Next:

	fmt.Printf("time: %#v \n", time.Now().Format(time.RFC1123))

	if parseDateTime, error := time.Parse(time.RFC3339, "2017-03-18T19:38:05Z"); error == nil {
		fmt.Printf("parse date: %v \n", parseDateTime)
	} else {
		fmt.Println("Error: " + error.Error())
	}

	if parseDateTime, error := time.Parse("2006-01-02", "2017-03-18"); error == nil {
		fmt.Printf("parse custom date based on time/format.go/const/std*** : %v \n", parseDateTime)
	} else {
		fmt.Println("Error: " + error.Error())
	}

	fmt.Printf("Using Date: %v \n", time.Date(2018, 3, 18, 19, 21, 00, 000, time.UTC))

	fmt.Println()
}
