package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Slide struct {
	Title string `xml:"title"`
	Type  string `xml:",type"`
}

type SlideShow struct {
	XMLName    xml.Name   `xml:"show"`
	TagComment string     `xml:",comment"`
	Author     string     `name: "field author" xml:"author,attr" json:"author, omitempty"`
	Date       string     `xml:"date,attr"`
	Attrs      []xml.Attr `xml:",attr"`

	// Slide      []Slide    `xml:"slide"` // using sub-element
	SlideTitle []string `xml:"slide>title"` // using xml path
}

func (s SlideShow) String() string {
	return fmt.Sprintf("Author:%v Date:%v TagComment: %v  Slides: %v ", s.Author, s.Date, s.TagComment, s.SlideTitle)
}

func Example() {
	file, _ := ioutil.TempFile("", "tempfile")
	defer func() {
		os.Remove(file.Name())
	}()

	var slideShow SlideShow = SlideShow{Author: "cherkavi", Date: "today", TagComment: "just a comment for a field", SlideTitle: []string{"one", "two", "three"}}
	fmt.Printf("> original data: %v\n", slideShow)

	// data, err := xml.Marshal(slideShow)
	// data, err := xml.MarshalIndent(slideShow, "", "  ")
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	err := encoder.Encode(slideShow)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println("> write data into file ")
	// ioutil.WriteFile(file.Name(), data, 666)

	fmt.Println("> read data from file ")
	text, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(text))

	var slideShowPerson SlideShow
	xml.Unmarshal(text, &slideShowPerson)
	fmt.Printf("> restored data: %v \n", slideShowPerson)
}
