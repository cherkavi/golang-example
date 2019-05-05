package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"log"
)

func main(){
	doc, err := goquery.NewDocument("http://ip-api.com/json")
  	if err != nil {
    	log.Fatal(err)
  	}
  	fmt.Print(dir(doc))
}
