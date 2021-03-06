package format

import (
	"fmt"
	"time"
)

func Example() {
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

	var nextTimePoint time.Time = time.Now().Add(2 * time.Hour)
	fmt.Printf("time point in 2 hours is: %v \n", nextTimePoint.Format(time.RFC1123))
	fmt.Println()
}
