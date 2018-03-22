package readget

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

func panicIfError(errorObject error, template string) {
	if errorObject != nil {
		panic(fmt.Sprintf(template, errorObject.Error()))
	}
}

func toUtf8(nonUtfData []byte) string {
	buf := make([]rune, len(nonUtfData))
	for i, b := range nonUtfData {
		buf[i] = rune(b)
	}
	return string(buf)
}

func readFromUrl(url string) []byte {
	response, error := http.Get(url)
	panicIfError(error, "can't read from remote source %v")
	defer response.Body.Close()

	content, error := ioutil.ReadAll(response.Body)
	panicIfError(error, "can't read from Http body: %v")
	return content
}

func readJsonFromUrl(url string) map[string]interface{} {
	var data map[string]interface{}
	if error := json.Unmarshal(readFromUrl(url), &data); error != nil {
		panic("can't read json data from text")
	}
	return data
}

func readXmlFromUrl(url string) SlideShow {
	data := readFromUrl(url)
	dataReader := bytes.NewReader(data)

	decoder := xml.NewDecoder(dataReader)
	decoder.CharsetReader = charset.NewReaderLabel

	var slideShow = SlideShow{}
	error := decoder.Decode(&slideShow)
	if error != nil {
		panic(fmt.Sprintf("can't create reader %v", error.Error()))
	}
	return slideShow
}

type Slide struct {
	Title string `xml:"title"`
	Type  string `xml:",type"`
}

type SlideShow struct {
	XMLName xml.Name   `xml:"slideshow"`
	Author  string     `xml:"author,attr"`
	Date    string     `xml:"date,attr"`
	Attrs   []xml.Attr `xml:",attr"`

	// Slide      []Slide    `xml:"slide"` // using sub-element
	SlideTitle []string `xml:"slide>title"` // using xml path
}

func Example() {
	fmt.Println("---------- read data from url ----------")
	fmt.Println(string(readFromUrl("http://ip.jsontest.com")))

	fmt.Println("---------- read json from url ----------")
	fmt.Println(readJsonFromUrl("http://validate.jsontest.com/?json={%22key%22:%22value%22}")["parse_time_nanoseconds"])

	fmt.Println("---------- read XML from url ----------")
	slideshow := readXmlFromUrl("https://httpbin.org/xml")
	fmt.Println(slideshow)
}
