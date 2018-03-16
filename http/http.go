package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


func panicIfError(error interface{}, template string) {
	if error != nil {
		panic(fmt.Sprintf(template, error))
	}
}

func readTextFromUrl(url string) []byte{
	response, error:=http.Get(url)
	panicIfError(error,"can't read from remote source %v")
	defer response.Body.Close()

	content, error :=ioutil.ReadAll(response.Body)
	panicIfError(error, "can't read from Http body: %v")
	return content
}

func readJsonFromUrl(url string) map[string]interface{}{
	var data map[string]interface{}
	if error:=json.Unmarshal(readTextFromUrl(url), &data); error!=nil{
		panic("can't read json data from text")
	}
	return data
}


func main(){
	fmt.Println(string(readTextFromUrl("http://ip.jsontest.com")))
	fmt.Println(readJsonFromUrl("http://validate.jsontest.com/?json={%22key%22:%22value%22}")["parse_time_nanoseconds"])
}